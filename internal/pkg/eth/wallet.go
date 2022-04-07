package eth

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/sodiumlabs/wallet-server/internal/pkg/types"
	"github.com/sodiumlabs/wallet-server/pkg/contracts"
	"github.com/sodiumlabs/wallet-server/pkg/rpcm"
)

var WalletFactoryAddressByNetwork = map[rpcm.Network]string{
	rpcm.Network_mumbai: "0x8b088C8243C207d1932D9C50fAcE12F8E5272A80",
}

func GetAddressForCounterfactualWallet(userId uint) (string, error) {
	client, err := rpcm.GetEthClient(rpcm.Network_mumbai)

	if err != nil {
		return "", err
	}

	factory, err := contracts.NewWalletFactory(ethcommon.HexToAddress(WalletFactoryAddressByNetwork[rpcm.Network_mumbai]), client)

	if err != nil {
		return "", err
	}

	b := types.UintToBytes20(userId)

	address, err := factory.GetAddressForCounterfactualWallet(
		&bind.CallOpts{},
		[]ethcommon.Address{ethcommon.HexToAddress("0x6544e298b1C1E8Dd13045ADa7B5d04b9238D77B4")},
		b,
	)

	if err != nil {
		return "", err
	}

	return address.Hex(), nil
}
