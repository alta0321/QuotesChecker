package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func setUpDatabase() {
	dbUrl := "user=postgres password=mysecretpassword dbname=postgres host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	err = DB.AutoMigrate(&Quote{})
	if err != nil {
		log.Fatal(err)
	}
}
