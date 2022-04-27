package routes

import (
	"TamaskaPJU/module/handler"
	"TamaskaPJU/module/utilities/device"
	"TamaskaPJU/module/utilities/user"
	"TamaskaPJU/module/view"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	// Load Repository
	userRepository := user.NewRepository(db)
	deviceRepository := device.NewRepository(db)
	//Load Service
	userService := user.NewService(userRepository)
	deviceService := device.NewService(deviceRepository)
	//Load Handler
	userHandler := handler.NewUserHandler(userService)
	//Load View
	deviceView := view.NewDeviceView(deviceService)

	// Routing Website Service
	router.GET("/device", deviceView.Index)
	router.GET("/device/edit/:id", deviceView.Edit)

	//Routing API Service
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)

	return router
}
