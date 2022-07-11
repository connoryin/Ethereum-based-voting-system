package data_access

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	cfg := mysql.Config{
		User:                 "admin",
		Passwd:               "cs6675project",
		Net:                  "tcp",
		Addr:                 "database-2.cisczib34pnn.us-east-1.rds.amazonaws.com:3306",
		DBName:               "sys",
		AllowNativePasswords: true,
	}

	var err error
	db, err = gorm.Open(gmysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected!")
}

func GetDB() *gorm.DB {
	return db
}
