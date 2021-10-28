package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName     string = "tctcore"
	Password     string = "tctcore"
	Addr         string = "192.168.68.105"
	Port         int    = 3306
	Database     string = "tctcore"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func CreateDatabase() *gorm.DB {
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		log.Println("Database connection Failed", err)
		return nil
	}
	return db
}
