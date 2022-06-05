package mpc

import (
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethereumhexutil "github.com/ethereum/go-ethereum/common/hexutil"
	ethereumcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/sodiumlabs/multi-party-sig/pkg/ecdsa"
	"github.com/sodiumlabs/multi-party-sig/pkg/math/curve"
	"github.com/sodiumlabs/multi-party-sig/pkg/party"
	"github.com/sodiumlabs/multi-party-sig/pkg/pool"
	"github.com/sodiumlabs/multi-party-sig/pkg/protocol"
	"github.com/sodiumlabs/multi-party-sig/protocols/cmp"
	"github.com/sodiumlabs/multi-party-sig/protocols/cmp/config"
	"github.com/sodiumlabs/wallet-server/internal/pkg/dao/model"
	"github.com/sodiumlabs/wallet-server/internal/test"
	"gorm.io/gorm"
)

var group = curve.Secp256k1{}

var configs = map[party.ID]*config.Config{}
var ids = party.IDSlice{}

func createTestMpcNode(db *gorm.DB) error {
	pl := pool.NewPool(0)
	threshold := 3
	configs, ids = test.GenerateConfig(group, 3, threshold, rand.Reader, pl)

	for _, id := range ids {
		configBytes, err := configs[id].MarshalBinary()

		if err != nil {
			return err
		}

		mpc := model.MPC{
			NodeId: string(id),
			Config: string(configBytes),
		}
		if err := db.Create(&mpc).Error; err != nil {
			return err
		}
	}

	return nil
}

func InitMPCNodeLocal(db *gorm.DB) error {
	var mpcNodes = []model.MPC{}
	if err := db.Find(&mpcNodes).Error; err != nil {
		return err
	}

	if len(mpcNodes) == 0 {
		return db.Transaction(createTestMpcNode)
	}

	for _, mpc := range mpcNodes {
		var config config.Config
		config.Group = group
		if err := config.UnmarshalBinary([]byte(mpc.Config)); err != nil {
			return err
		}
		ids = append(ids, config.ID)
		configs[party.ID(mpc.NodeId)] = &config

		// newConfig, err := config.DeriveBIP32(2)

		// if err != nil {
		// 	return err
		// }
	}

	if s, err := CMPSign(2, []byte("test")); err != nil {
		return err
	} else {
		b, _ := s.ToEthBytes()
		fmt.Println(ethereumhexutil.Encode(b))
	}

	return nil
}

func cmpSign(c *cmp.Config, m []byte, signers party.IDSlice, n *test.Network, pl *pool.Pool, sessionId []byte, cherror chan<- error, ch chan<- *ecdsa.Signature) {
	h, err := protocol.NewMultiHandler(cmp.Sign(c, signers, m, pl), sessionId)
	if err != nil {
		cherror <- err
		return
	}

	println("cmpSign:", string(c.ID))
	test.HandlerLoop(c.ID, h, n)

	if string(c.ID) == "a" {
		signResult, err := h.Result()

		if err != nil {
			cherror <- err
			return
		}

		signature := signResult.(*ecdsa.Signature)

		if !signature.Verify(c.PublicPoint(), m) {
			cherror <- errors.New("failed to verify cmp signature")
			return
		}

		ch <- signature

		return
	}
}

func CMPSign(userId uint, m []byte) (*ecdsa.Signature, error) {
	net := test.NewNetwork(ids)
	ch := make(chan *ecdsa.Signature)
	cherr := make(chan error)

	sessionId := ethereumcrypto.Keccak256(append(m, byte(userId)))

	for _, id := range ids {
		uconfigs, err := configs[id].DeriveBIP32(uint32(userId))

		if err != nil {
			return nil, err
		}

		go func(id party.ID) {
			pl := pool.NewPool(10)
			defer pl.TearDown()
			cmpSign(uconfigs, m, ids, net, pl, sessionId, cherr, ch)
		}(id)
	}

	select {
	case err := <-cherr:
		return nil, err
	case sig := <-ch:
		return sig, nil
	}
}

func GetUserAddress(userId uint) (*common.Address, error) {
	uconfigs, err := configs[ids[0]].DeriveBIP32(uint32(userId))

	if err != nil {
		return nil, err
	}

	uaddress := uconfigs.PublicPoint().ToAddress()

	return &uaddress, nil
}
