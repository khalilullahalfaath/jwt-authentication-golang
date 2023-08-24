package initializers

import (
	"fmt"

	"github.com/khalilullahalfaath/jwt-authentication-golang/models"
)

func SyncDatabase() {
	fmt.Println("Syncing database")
	err := DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("failed to sync database")
	}
}
