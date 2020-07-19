package db

import (
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/Projects/CabService_RestApi/models"
)

var db *gorm.DB
var err error

func Init() *gorm.DB {
	db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/cabrestapi")
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.Cab{}, &models.Customer{}, &models.Customerhistory{})
	return db
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
