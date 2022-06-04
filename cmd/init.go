package cmd

import (
	"os"

	"github.com/sodiumlabs/wallet-server/internal/pkg/db"
	"github.com/sodiumlabs/wallet-server/internal/pkg/mpc"
	"github.com/urfave/cli/v2"
)

/// 这里对一些服务进行初始化聚合
func xinit(c *cli.Context) error {
	inits := []func() error{
		// rpcm.Init,
	}

	dsn := os.Getenv("DB_DSN")

	inits = append(inits, func() error {
		return db.Init(dsn == "wallet.db", dsn)
	})

	inits = append(inits, func() error {
		return mpc.InitMPCNodeLocal(db.DB())
	})

	for _, init := range inits {
		if err := init(); err != nil {
			return err
		}
	}

	return nil
}
