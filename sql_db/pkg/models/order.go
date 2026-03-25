package models

import (
	"time"
)

type Order struct {
	ID         uint        // Standard field for the primary key
	UserID     uint        // Foreign key to Users table
	User       User        //
	Price      uint        // Price
	CreatedAt  time.Time   // Automatically managed by GORM for creation time
	OrderItems []OrderItem //
}
