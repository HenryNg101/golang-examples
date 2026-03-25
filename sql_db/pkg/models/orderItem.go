package models

import (
	"time"
)

type OrderItem struct {
	ID        uint      // Standard field for the primary key
	OrderID   uint      // Foreign key to Users table
	Order     Order     //
	Price     uint      // Price of the item
	Name      string    // Name of the item
	CreatedAt time.Time // Automatically managed by GORM for creation time
}
