package routes

import (
	"ClearningPatternGO/app/config"
	productv1 "ClearningPatternGO/modules/v1/utilities/product/view"
	userv1 "ClearningPatternGO/modules/v1/utilities/user/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	userHandlerV1 := userv1.Handler(db)
	productViewV1 := productv1.View(db)

	// Routing Website Service
	product := router.Group("/product")
	product.GET("/", productViewV1.Index)
	// router.GET("/device", deviceView.Index)
	// router.GET("/device/edit/:id", deviceView.Edit)

	//Routing API Service
	api := router.Group("/api/v1")
	api.POST("/users", userHandlerV1.RegisterUser)

	return router
}
