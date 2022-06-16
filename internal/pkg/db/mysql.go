package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func openMysql(dsn string) (*gorm.DB, error) {
	file, err := os.Create("/opt/app/gorm.log")

	if err != nil {
		return nil, err
	}

	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
}
