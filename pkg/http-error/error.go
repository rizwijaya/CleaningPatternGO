package error

import (
	"github.com/gin-gonic/gin"
)

func PageNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(404, "error_404.html", nil)
	}
}

func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"status": "404", "message": "Method Not Found"})
	}
}
