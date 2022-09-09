package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:rootPassword!@tcp(127.0.0.1:3306)/golang-bwa?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if (err != nil) {
		log.Fatal(err.Error())
	}

	log.Println("Successfully connected to database!")
}