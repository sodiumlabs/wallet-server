package ethapis

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/sodiumlabs/wallet-server/pkg/rpcm"
)

// data: "0x70a08231000000000000000000000000945e141f56ef168322798561503c4327e239da0e"
// from: undefined
// gas: undefined
// gasPrice: undefined
// to: "0x5a0585d409ca86d9fa771690ea37d32405da1f67"
type EthCallParams struct {
	Data     string  `json:"data"`
	From     *string `json:"from,omitempty"`
	Gas      *uint64 `json:"gas,omitempty"`
	GasPrice *string `json:"gasPrice,omitempty"`
	To       *string `json:"to,omitempty"`
}

func (p *EthCallParams) AsCallMsg() ethereum.CallMsg {
	msg := ethereum.CallMsg{
		From: ethcommon.BigToAddress(big.NewInt(0)),
		Data: common.FromHex(p.Data),
		Gas:  0,
	}

	if p.From != nil {
		msg.From = ethcommon.HexToAddress(*p.From)
	}

	if p.Gas != nil {
		msg.Gas = *p.Gas
	}

	if p.GasPrice != nil {
		msg.GasPrice = big.NewInt(0)
		msg.GasPrice.SetString(*p.GasPrice, 10)
	}

	if p.To != nil {
		to := ethcommon.HexToAddress(*p.To)
		msg.To = &to
	}

	return msg
}

type EthCallRequest struct {
	Id      int           `json:"id" validate:"required"`
	Jsonrpc string        `json:"jsonrpc" validate:"required"`
	Method  string        `json:"method" validate:"required"`
	Params  []interface{} `json:"params" validate:"required"`
}

type EthCallResponse struct {
	Result string `json:"result" validate:"required"`
}

// init wallet owner
func EthCall(c *gin.Context, q *EthCallRequest) (*EthCallResponse, error) {
	client, err := rpcm.GetEthClient(rpcm.Network_mumbai)

	if err != nil {
		return nil, err
	}

	if q.Method != "eth_call" {
		return nil, fmt.Errorf("invalid method: %s", q.Method)
	}

	if len(q.Params) != 2 {
		return nil, fmt.Errorf("invalid params: %v", q.Params)
	}

	params := EthCallParams{}
	var blocknumber *big.Int = nil

	if q.Params[1] != "latest" {
		blocknumber = big.NewInt(0)
		if _, ok := blocknumber.SetString(q.Params[1].(string), 0); !ok {
			return nil, fmt.Errorf("invalid blocknumber: %s", q.Params[2])
		}
	}

	j, err := json.Marshal(q.Params[0])

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(j, &params); err != nil {
		return nil, err
	}

	callmsg := params.AsCallMsg()

	// println(string(j))

	result, err := client.CallContract(context.TODO(), callmsg, blocknumber)

	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %s", err)
	}

	return &EthCallResponse{
		Result: fmt.Sprintf("0x%x", result),
	}, nil
}
