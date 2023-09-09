package database

import (
	"fmt"
	"go_blog/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DBConnection *gorm.DB

func ConnectionDB() {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")

	//dsn := "host=" + dbhost + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=America/Bogota"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota", dbhost, user, password, dbname, port)
	log.Println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	log.Println("Connection to database successfully")

	db.AutoMigrate(new(model.Blog))

	DBConnection = db

}
