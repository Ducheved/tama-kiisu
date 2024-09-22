package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Ducheved/tama-kiisu/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH environment variable is not set")
	}

	dbDir := filepath.Dir(dbPath)
	err := os.MkdirAll(dbDir, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create database directory: %v", err)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&models.Cat{})
	if err != nil {
		log.Fatalf("failed to perform migration: %v", err)
	}

	log.Println("Migration completed successfully")
}
