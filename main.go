package main

import (
	"ClearningPatternGO/app/config"
	"ClearningPatternGO/app/database"
	"ClearningPatternGO/modules/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	gin.SetMode(conf.App.Mode)
	db := database.Init(conf)
	router := routes.Init(db)

	//Load HTML Template
	router.Static("/assets", "./public/assets")
	router.Static("/images", "./public/images")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/fonts", "./public/assets/fonts")

	//fmt.Println("Starter " + conf.App.Name)
	router.Run(":" + conf.App.Port)
}
