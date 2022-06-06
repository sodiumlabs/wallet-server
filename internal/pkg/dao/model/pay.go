package model

import (
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/sodiumlabs/wallet-server/pkg/contracts"
	"github.com/sodiumlabs/wallet-server/pkg/rpcm"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"gorm.io/gorm"
)

//
type StripePay struct {
	Id                string `gorm:"primaryKey"`
	Amount            uint64
	UserWalletAddress string
	Status            stripe.CheckoutSessionStatus
	LatestCheckedAt   time.Time
}

// 更新支付状态
func StripePaySyncPool(db *gorm.DB) error {
	client, err := rpcm.GetEthClient(rpcm.Network_mumbai)

	if err != nil {
		return err
	}

	pkv := os.Getenv("PRIVATE_KEY")

	kp, err := secp256k1.NewKeypairFromString(pkv)

	if err != nil {
		return err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(kp.PrivateKey(), big.NewInt(80001))

	if err != nil {
		return err
	}

	opts.GasPrice = big.NewInt(55000000000)

	erc20, err := contracts.NewErc20(ethcommon.HexToAddress("0x5A0585D409ca86d9Fa771690ea37d32405Da1f67"), client)

	if err != nil {
		return err
	}

	for {
		var sp StripePay
		if err := db.Order("latest_checked_at asc").Where("status = ?", stripe.CheckoutSessionStatusOpen).First(&sp).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				fmt.Println(err.Error())
			}
			time.Sleep(time.Second * 60)
			continue
		}

		s, err := session.Get(sp.Id, nil)

		if err != nil {
			fmt.Println(err.Error(), "get session")
			time.Sleep(time.Second * 60)
			continue
		}

		sp.LatestCheckedAt = time.Now().UTC()
		sp.Status = s.Status

		if err := db.Save(&sp).Error; err != nil {
			fmt.Println(err.Error())
			time.Sleep(time.Second * 60)
			continue
		}

		if sp.Status == stripe.CheckoutSessionStatusComplete {
			amount := big.NewInt(0).Mul(big.NewInt(int64(sp.Amount)), big.NewInt(1e18))
			erc20.Transfer(opts, ethcommon.HexToAddress(sp.UserWalletAddress), amount)
		}
	}
}
