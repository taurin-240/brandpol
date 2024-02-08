package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func connectToPostgreSQL() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%v port=%s sslmode=disable",
		POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB_NAME, POSTGRES_HOST, POSTGRES_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checkErr(err)
	err = db.AutoMigrate(&Greet{})
	checkErr(err)
	return db
}

func get_history(db *gorm.DB) []Greet {
	var greet []Greet
	db.Find(&greet)
	return greet
}

func save_greet(db *gorm.DB, name string) error {
	var uuid = uuid.New()
	var createdAt = time.Now()
	var greet = fmt.Sprintf("Привет, %s от Go!", name)
	result := db.Create(&Greet{uuid, name, createdAt, greet})
	if result.Error != nil {
		return result.Error
	}
	return nil

}
