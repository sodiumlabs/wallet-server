package apis

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sodiumlabs/wallet-server/internal/pkg/dao/model"
	"github.com/sodiumlabs/wallet-server/internal/pkg/db"
	"github.com/sodiumlabs/wallet-server/internal/pkg/eth"
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

type InitWalletOwnerRequest struct {
	OwnerAddress             string `json:"owner_address" validate:"required"`
	OwnerEncryptedPrivatekey string `json:"owner_encrypted_privatekey" validate:"required"`
	Signature                string `json:"signature" validate:"required"`
}

type InitWalletOwnerResponse struct {
}

// init wallet owner
func InitWalletOwner(c *gin.Context, q *InitWalletOwnerRequest) (*InitWalletOwnerResponse, error) {
	userId := c.GetUint("user_id")

	var user model.User
	if err := db.DB().First(&user, userId).Error; err != nil {
		return nil, err
	}

	// Prevent repeated initialization, and restore the process if the owner needs to be changed
	if user.OwnerAddress != "" {
		return nil, fmt.Errorf("user already has owner address")
	}

	message := fmt.Sprintf("confirm %d", user.Id)

	signatureAddress, err := eth.VerifyMessage(message, q.Signature)

	if err != nil {
		return nil, err
	}

	if !strings.EqualFold(signatureAddress, q.OwnerAddress) {
		return nil, fmt.Errorf("signature address does not match %s %s", signatureAddress, q.OwnerAddress)
	}

	user.OwnerAddress = q.OwnerAddress
	user.OwnerEncryptedPrivatekey = q.OwnerEncryptedPrivatekey

	if err := db.DB().Save(&user).Error; err != nil {
		return nil, err
	}

	// create wallet deploy
	if err = db.DB().Transaction(func(tx *gorm.DB) error {
		for _, network := range model.FreeCreateWalletNetworks {
			if err := user.CreateWalletDeployNetworkForFree(tx, network); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &InitWalletOwnerResponse{}, nil
}
