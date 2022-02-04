package repository

import (
	"fmt"

	"github.com/absoran/goproject/models"
	"github.com/absoran/goproject/shared"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/postgres"
)

var gormdb *gorm.DB

func init() {
	db, err := gorm.Open("postgres", shared.Config.POSTGRESURL)
	if err != nil {
		panic(err.Error())
	}
	dbase := db.DB()
	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connection to database with gorm established.")
	}
	db.AutoMigrate(&models.User{})
	gormdb = db
}
