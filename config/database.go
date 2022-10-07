package config

import (
	"Go/WALIAPP-BACKEND/entity"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, dbname, port)

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.Teacher{}, &entity.Admin{})

	log.Println("connected on:", host, port)

	DB = db

	return db, nil
}

func ReadEnv() {
	viper.SetConfigName("App")
	viper.SetConfigType("yaml")
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("GoDotEnv:", &err)
	}
}