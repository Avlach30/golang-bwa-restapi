package auth

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetUsers() {
	dsn := "root:rootPassword!@tcp(127.0.0.1:3306)/golang-bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //* Configure gorm

	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("Connected to database!")

		var users []User //* Assign users slice with element of User struct
		
		db.Find(&users) //* MySql query 'select all from users' from users slice with implement referencing pointer

		for _, user := range users {
			fmt.Println(user.Name)
			fmt.Println(user.Email)
		}
	}
}