package handlers

import (
	"fmt"

	"github.com/noinch-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database connection driver
var database *gorm.DB

// InitDatabase to the database exported
func InitDatabase() error {
	var err error
	dsn := "root:root@tcp(localhost:5555)/noinch?parseTime=true"
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database!🔥")
	database.AutoMigrate(&models.User{}, &models.Manager{})
	fmt.Println("Database Migrated!🧨")
	return nil
}
