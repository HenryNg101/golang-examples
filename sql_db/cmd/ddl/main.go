package main

import (
	"log"

	"github.com/HenryNg101/golang-examples/sql_db/pkg"
	"github.com/HenryNg101/golang-examples/sql_db/pkg/db"
	"github.com/HenryNg101/golang-examples/sql_db/pkg/models"
	"gorm.io/gorm"
)

// Auto migration
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
	pkg.BulkCreate(db)
	// The one below will drop all tables and delete all data
	// dropTables(db)
}
