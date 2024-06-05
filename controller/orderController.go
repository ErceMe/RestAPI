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

	if err := ctx.ShouldBindJSON(&orders); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	config.GetDB().Create(&orders).Preload("Items")

	ctx.JSON(http.StatusOK, gin.H{"orders": orders, "Items": items})
}

func GetAllOrder(ctx *gin.Context) {
	var orders []entities.Order

	config.GetDB().Preload("Items").Find(&orders)
	ctx.JSON(http.StatusOK, gin.H{"orders": orders})
}

func GetOrderByID(ctx *gin.Context) {
	var order entities.Order
	id := ctx.Param("id")

	if err := config.GetDB().Preload("Items").First(&order, id).Error; err != nil {
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
	config.GetDB().Model(&orders).Association("Items")
	err := config.GetDB().Where("id = ?", id).Delete(&orders).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Failed to Delete"})
		fmt.Println("Error deleting data", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted Data!!!"})
}
