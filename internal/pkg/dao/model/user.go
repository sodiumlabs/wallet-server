package model

import (
	"fmt"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/sodiumlabs/wallet-server/internal/pkg/eth"
	"gorm.io/gorm"
)

type User struct {
	Id            uint   `gorm:"primaryKey;autoIncrement"`
	Username      string `gorm:"type:varchar(170);uniqueIndex"`
	WalletAddress string `gorm:"type:varchar(170);index"`
	Email         string `gorm:"type:varchar(170);index"`

	// current wallet owner
	OwnerAddress string `gorm:"type:varchar(170);index"`
	// current wallet owner encrypted private key
	// The user holds an encrypted private key, which ensures that the server has no control over any of the user's assets
	OwnerEncryptedPrivatekey string

	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
}

func (u *User) RWalletAddress() ethcommon.Address {
	return ethcommon.HexToAddress(u.WalletAddress)
}

func (u *User) ROwnerAddress() ethcommon.Address {
	return ethcommon.HexToAddress(u.OwnerAddress)
}

func (u *User) WalletIsDeployed(db *gorm.DB, network EthereumNetwork) (bool, error) {
	s, err := u.WalletDeployedStatus(db, network)

	if err != nil {
		return false, err
	}

	return s == WalletDeployStatusSuccess, nil
}

func (u *User) WalletDeployedStatus(db *gorm.DB, network EthereumNetwork) (WalletDeployStatus, error) {
	var walletd WalletDeploy
	if err := db.Find(&walletd, "user_id = ? and network = ?", u.Id, network).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return WalletDeployStatusNone, nil
		}
		return "", err
	}
	return walletd.Status, nil
}

func (u *User) OwnerBinded() bool {
	return u.OwnerAddress != "" &&
		u.OwnerEncryptedPrivatekey != ""
}

// Create a wallet for free
func (u *User) CreateWalletDeployNetworkForFree(db *gorm.DB, network EthereumNetwork) error {
	wd := &WalletDeploy{
		UserId:  u.Id,
		Network: network,
		Status:  WalletDeployStatusPending,

		// TODO
		ManagerSignature: "",
		OwnerSignature:   "",
		RefundAmount:     0,
		RefundToken:      "",
	}
	return db.Create(wd).Error
}

func RegisteUserByTw(db *gorm.DB, tuser *twitter.User) (*User, *UserTw, error) {
	if db == nil {
		return nil, nil, fmt.Errorf("db is nil")
	}
	if tuser == nil {
		return nil, nil, fmt.Errorf("tuser is nil")
	}

	var user *User
	var userTw *UserTw

	err := db.Transaction(func(tx *gorm.DB) error {
		user = &User{
			Username:                 fmt.Sprintf("tw_%d", tuser.ID),
			Email:                    tuser.Email,
			WalletAddress:            "",
			OwnerAddress:             "",
			OwnerEncryptedPrivatekey: "",
			CreatedAt:                time.Now().UTC(),
			UpdatedAt:                time.Now().UTC(),
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		userTw = &UserTw{
			TwId:   tuser.ID,
			UserId: user.Id,
		}
		if err := tx.Create(userTw).Error; err != nil {
			return err
		}

		// Wallet address calculation
		walletAddress, err := eth.GetAddressForCounterfactualWallet(user.Id)

		if err != nil {
			return err
		}

		user.WalletAddress = walletAddress

		if err := tx.Save(user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return user, userTw, nil
}

func RegisteUserByPhone(db *gorm.DB, phone string) (*User, *UserPhone, error) {
	if db == nil {
		return nil, nil, fmt.Errorf("db is nil")
	}

	var user *User
	var userPhone *UserPhone

	err := db.Transaction(func(tx *gorm.DB) error {
		user = &User{
			Username:                 fmt.Sprintf("phone_%d", phone),
			Email:                    "",
			WalletAddress:            "",
			OwnerAddress:             "",
			OwnerEncryptedPrivatekey: "",
			CreatedAt:                time.Now().UTC(),
			UpdatedAt:                time.Now().UTC(),
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		userPhone = &UserPhone{
			Phone:  phone,
			UserId: user.Id,
		}
		if err := tx.Create(userPhone).Error; err != nil {
			return err
		}

		// Wallet address calculation
		walletAddress, err := eth.GetAddressForCounterfactualWallet(user.Id)

		if err != nil {
			return err
		}

		user.WalletAddress = walletAddress

		if err := tx.Save(user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return user, userPhone, nil
}
