package main

import (
	"TamaskaPJU/app/auth"
	"TamaskaPJU/app/config"
	"TamaskaPJU/app/helper"
	"TamaskaPJU/module/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {
	db := config.Init()
	router := routes.Init(db)
	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("tamaskapju", cookieStore))

	//Load HTML Template
	router.HTMLRender = helper.Render("./public")
	router.Static("/assets", "./public/assets")
	router.Static("/images", "./public/images")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/fonts", "./public/assets/fonts")

	router.Run(":9000")
}
