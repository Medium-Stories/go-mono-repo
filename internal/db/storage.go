package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

// PostgresDb will initialize pg gorm database
func PostgresDb(connString string) *gorm.DB {
	if strings.TrimSpace(connString) == "" {
		logrus.Error("missing connection string")
		return nil
	}

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Error("failed to connect database: ", err)
		return nil
	}

	return db
}
