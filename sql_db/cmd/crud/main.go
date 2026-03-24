package main

import (
	"github.com/HenryNg101/golang-examples/sql_db/db"
	"github.com/HenryNg101/golang-examples/sql_db/pkg"
)

func main() {
	// var users []models.User
	db := db.ConnectDB()
	// db.Where("id < ?", "50").Find(&users)
	// fmt.Println(users)
	pkg.BulkCreate(db)
}
