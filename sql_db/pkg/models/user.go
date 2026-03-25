package models

import "time"

type User struct {
	ID        uint      // Standard field for the primary key
	Name      string    // A regular string field
	CreatedAt time.Time // Automatically managed by GORM for creation time
	Orders    []Order
}
