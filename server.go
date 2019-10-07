package main

import (
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func connect() *gorm.DB {
	db, err := gorm.Open("mysql", "ubuntu:ubuntu@tcp(127.0.0.1:3306)/workshifts?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to bbdd successfully!")
	return db
}
