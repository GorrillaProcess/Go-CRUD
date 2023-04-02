package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	//for some reason it needs to load env variables to migrate
	LoadEnvVars()

	var err error
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// fmt.Printf("DB string: %v", dsn)
	fmt.Printf("DB string: %v", DB)
	if err != nil {
		log.Fatal("Failed to connect to DB!")
	}
}
