package database

import (
	"webapp/server/config"
	"webapp/server/models"

	"github.com/golang/glog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	pgDB *gorm.DB
)

func Init() error {
	cfg := config.LoadConfig()
	dbPath := cfg.DatabasePath

	gormConfig := gorm.Config{}
	glog.Infof("Connecting to db: %s", dbPath)
	db, err := gorm.Open(postgres.Open(dbPath), &gormConfig)
	if err != nil {
		glog.Error(err)
		return err
	}
	glog.Info("Success")

	pgDB = db
	pgDB.AutoMigrate(
		&models.User{},
		&models.Account{},
		/*
			&models.Category{},
			&models.Merchant{},
			&models.Transaction{},
			&models.Rule{},
		*/
	)

	return nil
}

func GetDB() *gorm.DB {
	return pgDB
}
