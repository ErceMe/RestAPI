package main

import (
	"REST_API/config"
	"REST_API/router"
)

func main() {
	config.DBConnect()
	PORT := ":8080"
	router.StartServer().Run(PORT)
}
