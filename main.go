package main

import (
	"TamaskaPJU/app/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.Init()
	fmt.Println(db)
	router := gin.Default()

	router.Run(":9000")
}
