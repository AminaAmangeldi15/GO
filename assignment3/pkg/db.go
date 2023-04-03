package pkg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
