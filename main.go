package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
}

func DB() *gorm.DB {
	dsn := "host=localhost user=postgres password=9792amina dbname=golang sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB!")

	} else {
		fmt.Println("Connected to DB!")
	}
	return db
}

func main() {
	db := DB()
	db.Create(&Book{Id: "1", Title: "Fun", Description: "Enjoyyy", Cost: 1500})
}
