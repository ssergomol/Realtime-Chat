package database

import (
	"fmt"
	"log"

	"github.com/ssergomol/RealtimeChat/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost dbname=realtimechat"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfuly connected to database\n")

	DB = dbConn

	dbConn.AutoMigrate(&models.User{})

}
