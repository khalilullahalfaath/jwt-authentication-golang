package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global variable to hold the db connection
var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := `host=` + os.Getenv("DB_HOST") + ` user=` + os.Getenv("DB_USER") + ` password=` + os.Getenv("DB_PASSWORD") + ` dbname=` + os.Getenv("DB_NAME") + ` port=` + os.Getenv("DB_PORT") + ` sslmode=disable`

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
