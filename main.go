package main

import "REST_API/router"

func main() {
	PORT := "8080"

	router.StartServer().Run(PORT)
}
