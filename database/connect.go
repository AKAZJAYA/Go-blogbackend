package database

import (
	"log"
	"os"

	"github.com/AKAZJAYA/blogbackend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){

	err:=godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn:=os.Getenv("DSN")
	database, err:=gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{

		panic("Could not connect to database")
	}else {
		log.Print("Connect Successfully")
	}

	DB=database

	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}