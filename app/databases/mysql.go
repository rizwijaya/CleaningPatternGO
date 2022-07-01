package database

import (
	"ClearningPatternGO/app/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init(conf config.Conf) *gorm.DB {
	dsn := conf.Db.User + ":" + conf.Db.Pass + "@tcp(" + conf.Db.Host + ":" + conf.Db.Port + ")/" + conf.Db.Name + "?charset=utf8&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	return Db
}
