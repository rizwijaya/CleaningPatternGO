package routes

import (
	"ClearningPatternGO/app/config"
	productHandlerV1 "ClearningPatternGO/modules/v1/utilities/product/handler"
	productviewV1 "ClearningPatternGO/modules/v1/utilities/product/view"
	basic "ClearningPatternGO/pkg/basic_auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParseHTML(router *gin.Engine) *gin.Engine { //Load HTML Template
	router.Static("/assets", "./public/assets")
	router.Static("/images", "./public/images")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/fonts", "./public/assets/fonts")
	return router
}

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	productHandlerV1 := productHandlerV1.Handler(db)
	productViewV1 := productviewV1.View(db)

	// Routing Website Service
	product := router.Group("/product", basic.Auth(conf))
	product.GET("/", productViewV1.Index)

	//Routing API Service
	api := router.Group("/api/v1")
	api.GET("/product", productHandlerV1.ListProduct)

	router = ParseHTML(router)
	return router
}
