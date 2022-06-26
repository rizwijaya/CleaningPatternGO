package database

import (
	"log"

	config "ClearningPatternGO/app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(conf config.Conf) *gorm.DB {
	dsn := conf.Db.User + ":" + conf.Db.Pass + "@tcp(" + conf.Db.Host + ":" + conf.Db.Port + ")/" + conf.Db.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
