package config

import (
	"REST_API/entities"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func DBConnect() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading /env file")
	}

	host := os.Getenv("dbHost")
	user := os.Getenv("dbUser")
	password := os.Getenv("dbPassword")
	dbport := os.Getenv("dbPort")
	dbname := os.Getenv("dbName")

	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbport, dbname)
	fmt.Println(config)
	db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	db.AutoMigrate(&entities.Order{}, &entities.Item{})
}

func GetDB() *gorm.DB {
	return db
}
