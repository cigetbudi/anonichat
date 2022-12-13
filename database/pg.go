package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("gagal membaca file .env ", err)
	}
	return os.Getenv(key)
}

func InitDB() {
	var (
		host     = getEnv("DB_HOST")
		port     = getEnv("DB_PORT")
		user     = getEnv("DB_USER")
		dbname   = getEnv("DB_NAME")
		password = getEnv("DB_PASSWORD")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate()
	fmt.Println("sukses terhubung dengan db")
}
