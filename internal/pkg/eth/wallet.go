package eth

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/sodiumlabs/wallet-server/internal/pkg/types"
	"github.com/sodiumlabs/wallet-server/pkg/contracts"
	"github.com/sodiumlabs/wallet-server/pkg/rpcm"
)

var WalletFactoryAddressByNetwork = map[rpcm.Network]string{
	rpcm.Network_mumbai: "0x0A02100e0350D89D1eBDF08d8728297B1a8f85eF",
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
		[]ethcommon.Address{ethcommon.HexToAddress("0x968aD39B0812e853685d8567fb8AA009E63F2fcc")},
		b,
	)

	if err != nil {
		return "", err
	}

	return address.Hex(), nil
}
