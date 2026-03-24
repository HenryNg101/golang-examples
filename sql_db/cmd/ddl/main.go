package main

import (
	"log"

	"github.com/HenryNg101/golang-examples/sql_db/db"
	"github.com/HenryNg101/golang-examples/sql_db/models"
	"gorm.io/gorm"
)

func createOrAlterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("Migration error: ", err)
	}
}

func dropTables(db *gorm.DB) {
	err := db.Migrator().DropTable(
		&models.User{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("Drop tables error: ", err)
	}
}

func main() {
	db := db.ConnectDB()
	createOrAlterTables(db)
}
