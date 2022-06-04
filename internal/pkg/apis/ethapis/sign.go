package ethapis

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/sodiumlabs/wallet-server/internal/pkg/mpc"
)

type EthSignParams struct {
	Message     string `json:"message"`
	MessageHash string `json:"messageHash"`
}

type EthSignResponse struct {
	Signature string `json:"signature" validate:"required"`
}

// init wallet owner
func EthSign(c *gin.Context, q *EthSignParams) (*EthSignResponse, error) {
	userId := c.GetUint("user_id")

	sig, err := mpc.CMPSign(userId, []byte(q.MessageHash))

	if err != nil {
		return nil, err
	}

	signBytes, err := sig.ToEthBytes()

	if err != nil {
		return nil, err
	}

	sig.R.MarshalBinary()

	res := EthSignResponse{
		Signature: common.Bytes2Hex(signBytes),
	}

	return &res, nil
}
