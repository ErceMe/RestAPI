package controller

import (
	"REST_API/config"
	"REST_API/entities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(ctx *gin.Context) {
	var orders entities.Order
	var items entities.Item

	fmt.Println(orders)
	if err := ctx.ShouldBindJSON(&orders); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err := config.GetDB().Create(&orders).Preload("items")
	if err != nil {
		fmt.Println("Error creating Data!!!!")
	}
	ctx.JSON(http.StatusOK, gin.H{"orders": orders, "items": items})
}

func GetAllOrder(ctx *gin.Context) {
	var orders []entities.Order

	config.GetDB().Preload("items").Find(&orders)
	ctx.JSON(http.StatusOK, gin.H{"orders": orders})
}

func GetOrderByID(ctx *gin.Context) {
	var order entities.Order
	id := ctx.Param("id")

	if err := config.GetDB().Preload("items").First(&order, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"order": order})
}

func UpdateOrder(ctx *gin.Context) {
	var orders entities.Order
	// var items entities.Item

	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&orders); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	if config.GetDB().Model(&orders).Where("id = ?", id).Updates(&orders).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"orders": orders})
}

func DeleteOrder(ctx *gin.Context) {
	var orders entities.Order
	id := ctx.Param("id")
	err := config.GetDB().Where("id = ?", id).Delete(&orders).Error
	if err != nil {
		fmt.Println("Error deleting data", err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted Data!!!"})
}
