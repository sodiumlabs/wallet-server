package model

import (
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/sodiumlabs/wallet-server/internal/pkg/eth"
	"github.com/sodiumlabs/wallet-server/internal/pkg/types"
	"github.com/sodiumlabs/wallet-server/pkg/contracts"
	"github.com/sodiumlabs/wallet-server/pkg/rpcm"
	"gorm.io/gorm"
)

type EthereumNetwork string

const (
	EthereumNetworkMainnet = EthereumNetwork("mainnet")
	EthereumNetworkMumbai  = EthereumNetwork("mumbai")
	EthereumNetworkMatic   = EthereumNetwork("matic")
)

var FreeCreateWalletNetworks = []EthereumNetwork{
	EthereumNetworkMumbai,
}

var SupportedNetworks = []EthereumNetwork{
	EthereumNetworkMumbai,
	EthereumNetworkMatic,
}

type WalletDeployStatus = string

const (
	WalletDeployStatusNone    = WalletDeployStatus("none")
	WalletDeployStatusPending = WalletDeployStatus("pending")
	WalletDeployStatusSuccess = WalletDeployStatus("success")

	// Wallet deployment failure, which occurs when it costs more to deploy some blocks and the user's wallet does not have enough gas fee token.
	WalletDeployStatusFailed = WalletDeployStatus("failed")
)

// Wallet represents a user's wallet.
type WalletDeploy struct {
	Id               uint            `gorm:"primaryKey;autoIncrement"`
	UserId           uint            `gorm:"index;uniqueIndex:userNetworkIDX"`
	Network          EthereumNetwork `gorm:"type:varchar(170);uniqueIndex:userNetworkIDX"`
	TxHash           string
	Status           WalletDeployStatus
	ManagerSignature string
	OwnerSignature   string
	RefundAmount     uint64
	RefundToken      string
}

func DefaultGuardian() ethcommon.Address {
	return ethcommon.HexToAddress("0xCe9e70007e6Fa8873f666A498f59451b7c27AF31")
}

func WalletModules() []ethcommon.Address {
	return []ethcommon.Address{
		ethcommon.HexToAddress("0x3436d1E60647A71407555e6ab5b62Bbe22D2a5b2"),
	}
}

func WalletDeployPool(db *gorm.DB) error {
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

	factory, err := contracts.NewWalletFactory(ethcommon.HexToAddress(eth.WalletFactoryAddressByNetwork[rpcm.Network_mumbai]), client)

	if err != nil {
		return err
	}

	for {
		var wd WalletDeploy
		if err := db.First(&wd, "status = ?", WalletDeployStatusPending).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				fmt.Println(err.Error())
				time.Sleep(60 * time.Second)
				continue
			}
		}

		var user User
		if err := db.First(&user, "id = ?", wd.UserId).Error; err != nil {
			fmt.Println(err.Error())
			time.Sleep(60 * time.Second)
			continue
		}

		// address _owner,
		// address[] calldata _modules,
		// address _guardian,
		// bytes20 _salt,
		// uint256 _refundAmount,
		// address _refundToken,
		// bytes calldata _ownerSignature,
		// bytes calldata _managerSignature
		t, err := factory.CreateCounterfactualWallet(
			opts,
			user.ROwnerAddress(),
			WalletModules(),
			DefaultGuardian(),
			types.UintToBytes20(wd.UserId),
			big.NewInt(0),
			ethcommon.HexToAddress("0x0000000000000000000000000000000000000000"),
			[]byte(wd.OwnerSignature),
			[]byte(wd.ManagerSignature),
		)

		if err != nil {
			fmt.Println(err.Error(), t, "wallet error")
			wd.Status = WalletDeployStatusSuccess
			db.Save(&wd)
			continue
		}

		ch := make(chan *contracts.WalletFactoryWalletCreated)

		sub, err := factory.WatchWalletCreated(
			&bind.WatchOpts{},
			ch,
			[]ethcommon.Address{user.RWalletAddress()},
			[]ethcommon.Address{user.ROwnerAddress()},
			[]ethcommon.Address{DefaultGuardian()},
		)

		if err != nil {
			fmt.Println(err.Error())
		}

		<-ch

		wd.TxHash = t.Hash().Hex()
		wd.Status = WalletDeployStatusSuccess
		db.Save(&wd)

		sub.Unsubscribe()
	}
}
