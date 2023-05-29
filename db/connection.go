package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DSN = "root:root@tcp(localhost:3306)/gorm?parseTime=true"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB connected")
	}
}
