package router

import (
	"REST_API/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controller.CreateOrder)
	router.GET("/orders", controller.GetAllOrder)
	router.GET("/orders/:id", controller.GetOrderByID)
	router.PUT("/orders/:id", controller.UpdateOrder)
	router.DELETE("/orders/:id", controller.DeleteOrder)

	return router
}
