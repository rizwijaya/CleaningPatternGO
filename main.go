package main

import (
	"ClearningPatternGO/app/config"
	database "ClearningPatternGO/app/databases"
	"ClearningPatternGO/app/helper"
	"ClearningPatternGO/modules/v1/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.Init()
	gin.SetMode(conf.App.Mode)
	if err != nil {
		log.Fatal(err)
	}
	db := database.Init(conf)

	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(conf.App.Secret_key))
	router.Use(sessions.Sessions("cleaningpatterngo", cookieStore))
	router.HTMLRender = helper.Render("./public")

	//Error Handling for 404 Not Found Page and Method Not Allowed
	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "error_404.html", nil)
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": "404", "message": "Method Not Found"})
	})
	router = routes.Init(db, conf, router)

	//Load HTML Template
	router.Static("/assets", "./public/assets")
	router.Static("/images", "./public/images")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/fonts", "./public/assets/fonts")

	//fmt.Println("Starter " + conf.App.Name)
	router.Run(":" + conf.App.Port)
}
