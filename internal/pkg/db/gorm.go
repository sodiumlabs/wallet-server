package db

import (
	"github.com/sodiumlabs/wallet-server/internal/pkg/dao/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func autoMigrate() error {
	models := []interface{}{
		model.User{},
		model.UserTw{},
		model.WalletDeploy{},
		model.StripePay{},
		model.UserPhone{},
	}
	return db.AutoMigrate(models...)
}

func Init(isDev bool, dsn string) error {
	var err error
	if isDev {
		db, err = openSqlite(dsn)
	} else {
		db, err = openMysql(dsn)
	}

	if err != nil {
		return err
	}

	err = autoMigrate()

	return err
}

func DB() *gorm.DB {
	if db == nil {
		panic("db uninitialized")
	}
	return db
}
