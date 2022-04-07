package ethapis

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/sodiumlabs/wallet-server/internal/pkg/dao/model"
	"github.com/sodiumlabs/wallet-server/internal/pkg/db"
	"github.com/sodiumlabs/wallet-server/pkg/contracts"
	"github.com/sodiumlabs/wallet-server/pkg/rpcm"
)

// data: "0x70a08231000000000000000000000000945e141f56ef168322798561503c4327e239da0e"
// from: undefined
// gas: undefined
// gasPrice: undefined
// to: "0x5a0585d409ca86d9fa771690ea37d32405da1f67"
type EthTxParams struct {
	Data      string  `json:"data"`
	From      *string `json:"from,omitempty"`
	Gas       *uint64 `json:"gas,omitempty"`
	GasPrice  *string `json:"gasPrice,omitempty"`
	Nonce     int64   `json:"nonce"`
	Signature string  `json:"signature"`
	To        string  `json:"to"`
}

// init wallet owner
func EthTx(c *gin.Context, q *EthCallRequest) (*EthCallResponse, error) {
	client, err := rpcm.GetEthClient(rpcm.Network_mumbai)

	userId := c.GetUint("user_id")

	var user model.User

	if err := db.DB().First(&user, userId).Error; err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if q.Method != "eth_sendTransaction" {
		return nil, fmt.Errorf("invalid method: %s", q.Method)
	}

	if len(q.Params) != 1 {
		return nil, fmt.Errorf("invalid params: %v", q.Params)
	}

	pkv := os.Getenv("PRIVATE_KEY")

	kp, err := secp256k1.NewKeypairFromString(pkv)

	if err != nil {
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(kp.PrivateKey(), big.NewInt(80001))
	opts2, err := bind.NewKeyedTransactorWithChainID(kp.PrivateKey(), big.NewInt(80001))

	if err != nil {
		return nil, err
	}

	opts.GasPrice = big.NewInt(55000000000)

	module, err := contracts.NewArgentModule(model.WalletModules()[0], client)

	if err != nil {
		return nil, err
	}

	params := EthTxParams{}

	j, err := json.Marshal(q.Params[0])

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(j, &params); err != nil {
		return nil, err
	}

	opts2.GasLimit = 21000
	opts2.NoSend = true
	rtx, err := module.MultiCallWithGuardians(opts2, user.RWalletAddress(), []contracts.TransactionManagerCall{
		{
			To:    ethcommon.HexToAddress(params.To),
			Value: big.NewInt(0),
			Data:  common.FromHex(params.Data),
		},
	})

	if err != nil {
		return nil, fmt.Errorf("ready failed: %s", err)
	}

	if err := json.Unmarshal(j, &params); err != nil {
		return nil, err
	}

	ss := []string{
		params.Signature,
	}

	jj, err := json.Marshal(ss)

	if err != nil {
		return nil, fmt.Errorf("ready failed: %s", err)
	}

	tx, err := module.Execute(
		opts,
		common.HexToAddress(user.WalletAddress),
		rtx.Data(),
		big.NewInt(params.Nonce),
		[]byte(jj),
		big.NewInt(0),
		big.NewInt(0),
		common.HexToAddress("0x5A0585D409ca86d9Fa771690ea37d32405Da1f67"),
		common.HexToAddress("0xf521F6d48703DE80Ba378cFf4Ea7519272d315B7"),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %s", err)
	}

	return &EthCallResponse{
		Result: tx.Hash().Hex(),
	}, nil
}
