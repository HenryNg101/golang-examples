package main

import (
	"fmt"
	"log"

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

	userCount, err := services.CountUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The current amount of users there are in the system is: ", userCount)

	// Aggregation test
	topRequired := 20
	topTenUsers, err := services.GetTopSpenders(db, topRequired)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Top %d users in spending:\n", topRequired)
	for _, user := range topTenUsers {
		fmt.Printf("User ID: %d. Total spending: %d\n", user.UserID, user.Total)
	}
	fmt.Println()

	// Nested preloading test
	userInfo, err := services.GetUserOrders(db, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User %s's orders: \n", userInfo.Name)
	for _, order := range userInfo.Orders {
		fmt.Println()
		fmt.Println("Order ID: ", order.ID)
		fmt.Println("Order items: ")
		for _, orderItem := range order.OrderItems {
			fmt.Println("-\tOrder item's name: ", orderItem.Name)
			fmt.Println("\tOrder item's price: ", orderItem.Price)
		}
	}

	// Subquery test
	orders, err := services.GetExpensiveOrders(db)
	for _, order := range orders {
		fmt.Println()
		fmt.Println("Order ID: ", order.ID)
		fmt.Println("Order price: ", order.Price)
	}
}
