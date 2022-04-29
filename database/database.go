package database

import (
	"fmt"
	"github.com/azifaazka/MiniProject-Novellist/config"
	"github.com/azifaazka/MiniProject-Novellist/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func InitDB(conf config.Config) {

	conectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(conectionString))
	if err != nil {
		fmt.Println("error open conection : ", err)
	}

	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Bab{})
	DB.AutoMigrate(&model.Star{})
	DB.AutoMigrate(&model.Cerita{})
	DB.AutoMigrate(&model.Comment{})
}