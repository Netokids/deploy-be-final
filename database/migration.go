package database

import (
	"BE-finaltask/models"
	"BE-finaltask/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Consultation{},
	)
	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}
	fmt.Println("Migration Success")
}
