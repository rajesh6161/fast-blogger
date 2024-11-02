// db/database.go
package db

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

// It ensures that the initialization code within the Do block is executed only once during the application’s lifetime, even if Initialize is called multiple times.
func Initialize(dsn string) *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
		instance = db
		log.Println("Database connection established")
	})
	return instance
}

// GetDB returns the singleton database instance
func GetDB() *gorm.DB {
	if instance == nil {
		log.Fatal("Database is not initialized. Please call db.Initialize() first.")
	}
	return instance
}
