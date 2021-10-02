package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB() error {
	// TODO: Refactor dsn
	dsn := "host=localhost user=postgres password=postgres dbname=tgtc_user_coupon port=3306 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}
