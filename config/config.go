package config

import (
	structs "assigment2/structs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit(dbUsername, dbPassword, dbPort, dbName, dbHost string) *gorm.DB {
	dsn := fmt.Sprint("root:12345@tcp(127.0.0.1:3306)/orders_by")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(structs.Order{})
	db.AutoMigrate(structs.Item{})
	return db
}
