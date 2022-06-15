package apis

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sodiumlabs/wallet-server/internal/pkg/dao/model"
	"github.com/sodiumlabs/wallet-server/internal/pkg/db"
	"github.com/sodiumlabs/wallet-server/internal/pkg/mpc"
	"gorm.io/gorm"
)

type WalletNetworkStatus struct {
	Network model.EthereumNetwork    `json:"network"`
	Status  model.WalletDeployStatus `json:"status"`
}

// 是否绑定
type WalletInfoResponse struct {
	UserId                   uint                  `json:"user_id" validate:"required"`
	WalletAddress            string                `json:"wallet_address" validate:"required"`
	OwnerAddress             string                `json:"owner_address" validate:"required"`
	OwnerEncryptedPrivatekey string                `json:"owner_encrypted_privatekey" validate:"required"`
	Networks                 []WalletNetworkStatus `json:"networks" validate:"required"`
}

func WalletInfo(c *gin.Context) (*WalletInfoResponse, error) {
	userId := c.GetUint("user_id")

	user := model.User{}
	if err := db.DB().First(&user, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(401)
		}
		return nil, err
	}

	networks := []WalletNetworkStatus{}

	for _, network := range model.SupportedNetworks {
		status, err := user.WalletDeployedStatus(db.DB(), network)
		if err != nil {
			return nil, err
		}

		networks = append(networks, WalletNetworkStatus{
			Network: network,
			Status:  status,
		})
	}

	return &WalletInfoResponse{
		UserId:                   userId,
		WalletAddress:            user.WalletAddress,
		OwnerAddress:             user.OwnerAddress,
		OwnerEncryptedPrivatekey: user.OwnerEncryptedPrivatekey,
		Networks:                 networks,
	}, nil
}

type SignLogRequest struct {
	Address string `json:"address" validate:"required"`
	Line    int32  `json:"line" validate:"required"`
}

type SignLogResponse struct {
	Contents   []string `json:"contents" validate:"required"`
	LatestLine int      `json:"latest_line" validate:"required"`
}

func SignLogApi(c *gin.Context, q *SignLogRequest) (*SignLogResponse, error) {
	filename := fmt.Sprintf("/opt/app/%s.log", strings.ToLower(q.Address))
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	contents := []string{}
	line := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		contents = append(contents, scanner.Text())
		line++
	}

	return &SignLogResponse{
		Contents:   contents,
		LatestLine: line,
	}, nil
}

type InitWalletOwnerRequest struct {
	OwnerAddress             string `json:"owner_address" validate:"required"`
	OwnerEncryptedPrivatekey string `json:"owner_encrypted_privatekey" validate:"required"`
	Signature                string `json:"signature" validate:"required"`
}

type InitWalletOwnerResponse struct {
}

// init wallet owner
func InitWalletOwner(userId uint, db *gorm.DB) error {
	var user model.User
	if err := db.First(&user, userId).Error; err != nil {
		return err
	}

	// Prevent repeated initialization, and restore the process if the owner needs to be changed
	if user.OwnerAddress != "" {
		return fmt.Errorf("user already has owner address")
	}

	address, err := mpc.GetUserAddress(userId)

	if err != nil {
		return err
	}

	user.OwnerAddress = address.Hex()
	user.OwnerEncryptedPrivatekey = ""

	if err := db.Save(&user).Error; err != nil {
		return err
	}

	// create wallet deploy
	if err = db.Transaction(func(tx *gorm.DB) error {
		for _, network := range model.FreeCreateWalletNetworks {
			if err := user.CreateWalletDeployNetworkForFree(tx, network); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
