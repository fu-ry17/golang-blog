package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"simple-blog/config"
	"simple-blog/db/models"
)

func NewStorage() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Env.DB_USER, config.Env.DB_PASSWORD, config.Env.DB_ADDRESS, config.Env.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.Logger = db.Logger.LogMode(logger.Info)
	fmt.Println("Connected to database...")
	// run migrations
	db.AutoMigrate(&models.Blog{})

	return db, nil
}
