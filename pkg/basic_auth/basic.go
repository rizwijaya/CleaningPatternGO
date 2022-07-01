package basic

import (
	"ClearningPatternGO/app/config"

	"github.com/gin-gonic/gin"
)

func Auth(conf config.Conf) gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		conf.BasicAuth.Username: conf.BasicAuth.Password,
	})
}

func AuthAPI(conf config.Conf) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		if user != conf.BasicAuth.Username || pass != conf.BasicAuth.Password {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
	}
}
