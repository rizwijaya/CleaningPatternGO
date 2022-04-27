package main

import (
	"TamaskaPJU/app/config"
	"TamaskaPJU/module/routes"
)

func main() {
	db := config.Init()
	router := routes.Init(db)

	router.Run(":9000")
}
