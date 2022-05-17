package persistance

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var err error

func DBconnection(dsn string) (*gorm.DB, error) {
	cfg := mysql.Config{
		DSN: dsn,
	}
	DB, err = gorm.Open(mysql.New(cfg), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return DB, nil
}
