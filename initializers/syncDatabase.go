package initializers

import (
	"accounting-project/models"
	"accounting-project/pkg/postgres"
	"fmt"
)

func SyncDatabase() {
	err := postgres.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Card{},
		&models.Bank{},
	)
	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
