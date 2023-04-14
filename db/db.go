package db

import (
	"log"

	"Employee-Log-System/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:Sriven@123@localhost/employee"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.User{}, &models.Department{}, &models.Employee{}, &models.Attendance{})

	return db
}
