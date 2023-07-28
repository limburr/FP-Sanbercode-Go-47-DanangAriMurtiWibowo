package config

import (
	"finalproj/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	username := "root"
	password := ""
	host := "tcp(127.0.0.1:3306)"
	database := "blog_db"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charsetutf8mb4&parseTime=True&loc=Local", username, password, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Comment{}, &models.Post{}, &models.Role{})

	return db
}
