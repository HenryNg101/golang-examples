package main

import (
	"fmt"
	"log"

	"github.com/HenryNg101/golang-examples/sql_db/pkg"
	"github.com/HenryNg101/golang-examples/sql_db/pkg/db"
	"github.com/HenryNg101/golang-examples/sql_db/pkg/services"
)

func main() {
	db := db.ConnectDB()
	// CRUD on users
	createdUser, err := services.CreateUser(db, "Andrew")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(createdUser)

	userCount, err := services.ListUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The current amount of users there are in the database are: ", userCount)

	// CRUD on orders

	// CRUD on order items

	pkg.BulkCreate(db)
}
