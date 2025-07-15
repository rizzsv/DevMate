package config

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Gunakan postgres.Config + PreferSimpleProtocol
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // ✅ matikan prepared statement dari driver
	}), &gorm.Config{
		PrepareStmt: false, // ✅ matikan statement cache GORM
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
}
