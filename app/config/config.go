package config

import (
	"fmt"
	"log"

	"github.com/tkanos/gonfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

var db *gorm.DB

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("app/config/config.json", &conf)
	return conf
}

func Init() *gorm.DB {
	conf := GetConfig()
	dsn := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	//dsn := "root:@tcp(127.0.0.1:3306)/crowfund?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println(db)
		log.Println("Connected to database")
	}
	return db
}
