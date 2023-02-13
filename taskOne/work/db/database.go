package db

import (
	"log"
	"os"
	"work/constants"
	"work/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connect to db
func Connect() *gorm.DB {
	dsn := os.Getenv(constants.DbDsn)
	//fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("%v", err)
		return nil
	}
	AutoMigrate(db)
	return db
}

// automatic migration db
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(model.Order{})
}
