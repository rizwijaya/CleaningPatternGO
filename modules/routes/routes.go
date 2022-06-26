package routes

import (
	"ClearningPatternGO/app/auth"
	"ClearningPatternGO/app/helper"
	"ClearningPatternGO/modules/handler"
	"ClearningPatternGO/modules/utilities/device"
	"ClearningPatternGO/modules/utilities/user"
	"ClearningPatternGO/modules/view"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gin.Engine {

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

	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("tamaskapju", cookieStore))
	router.HTMLRender = helper.Render("./public")

	// Routing Website Service
	router.GET("/device", deviceView.Index)
	router.GET("/device/edit/:id", deviceView.Edit)

	//Routing API Service
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)

	return router
}
