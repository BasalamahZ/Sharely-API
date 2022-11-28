package configs

import (
	"log"
	"sharely/models"
)

func SyncDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error Migrate")
	}
}


