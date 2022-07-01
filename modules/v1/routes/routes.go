package routes

import (
	"ClearningPatternGO/app/config"
	productHandlerV1 "ClearningPatternGO/modules/v1/utilities/product/handler"
	productviewV1 "ClearningPatternGO/modules/v1/utilities/product/view"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	//userHandlerV1 := userHandlerV1.Handler(db)
	productHandlerV1 := productHandlerV1.Handler(db)
	productViewV1 := productviewV1.View(db)

	// Routing Website Service
	product := router.Group("/product")
	product.GET("/", productViewV1.Index)

	//Routing API Service
	api := router.Group("/api/v1")
	api.GET("/product", productHandlerV1.ListProduct)
	//api.POST("/users", userHandlerV1.RegisterUser)

	return router
}
