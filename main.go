package main

import (
	"ClearningPatternGO/app/config"
	"ClearningPatternGO/app/database"
	"ClearningPatternGO/modules/routes"
	"log"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db := database.Init(conf)
	router := routes.Init(db)

	//Load HTML Template
	router.Static("/assets", "./public/assets")
	router.Static("/images", "./public/images")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/fonts", "./public/assets/fonts")

	router.Run(":9000")
}
