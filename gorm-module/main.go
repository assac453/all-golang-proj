package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to init db")
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{
		Code:  "D42",
		Price: 100,
	})

	var product Product

	db.First(&product, 1)
	fmt.Println(product.ID)

	db.First(&product, "code = ?", "D42")
	fmt.Println(product.ID)

}
