package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	//Routing API Service
	api := router.Group("/api/v1")

	return router
}
