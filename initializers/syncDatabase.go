package initializers

import (
	"crud/models"
	"fmt"
)

func SyncDatabase() {
	// Migrate users table
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		fmt.Println(err)
	}
	// Migrate posts table
	if err := DB.AutoMigrate(&models.Post{}); err != nil {
		fmt.Println(err)
	}
}
