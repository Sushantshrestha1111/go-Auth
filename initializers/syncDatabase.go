package initializers

import "go-auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

}
