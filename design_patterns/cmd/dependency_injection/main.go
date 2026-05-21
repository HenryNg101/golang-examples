package main

import "fmt"

// Interface for injection, just a good way to combine both DIP and Dependency Injection pattern
type Saver interface {
	Save(string)
}

type PostgresDatabase struct{}

func (db PostgresDatabase) Save(data string) {
	fmt.Println("Saving into Postgres...")
	fmt.Println("The username to be saved is:", data)
}

/*
You can even have more DB types here

type MySQLDatabase struct{}

func (db MySQLDatabase) Save(data string) {}

Then in main, just replace db := PostgresDatabase{} to db := MySQLDatabase{}, and it still works perfectly
*/

// The service doesn't have to worry of the DB initialization, it just uses it
type UserService struct {
	db Saver
}

func (s UserService) CreateUser(name string) {
	s.db.Save(name)
}

func main() {
	db := PostgresDatabase{}
	service := UserService{db: db} // Inject into the service, instead of letting the service having to create it inside the implementation
	service.CreateUser("Henry")
}
