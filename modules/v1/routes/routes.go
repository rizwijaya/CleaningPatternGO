package routes

import (
	"ClearningPatternGO/app/config"
	//productviewv1 "ClearningPatternGO/modules/v1/utilities/product/view"
	productHandlerV1 "ClearningPatternGO/modules/v1/utilities/product/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	//userHandlerV1 := userHandlerV1.Handler(db)
	productHandlerV1 := productHandlerV1.Handler(db)
	//productViewV1 := productviewv1.View(db)

	// Routing Website Service
	product := router.Group("/product")
	product.GET("/", productHandlerV1.ListProduct)

	//Routing API Service
	//api := router.Group("/api/v1")
	//api.POST("/users", userHandlerV1.RegisterUser)

	return router
}
