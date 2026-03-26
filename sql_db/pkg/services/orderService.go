package services

import (
	"github.com/HenryNg101/golang-examples/sql_db/pkg/models"

	"gorm.io/gorm"
)

// Create a transaction, in which it will try creating an order.
// If the anon function returns nil, transaction commits
// If the anon function return errors, transactions rollback
func CreateOrder(db *gorm.DB, userID uint, items []models.OrderItem) error {
	return db.Transaction(func(tx *gorm.DB) error {
		order := models.Order{
			UserID: userID,
			Price:  0,
		}

		if err := tx.Create(&order).Error; err != nil { // Create an order
			return err
		}

		var total uint = 0

		for i := range items {
			items[i].OrderID = order.ID
			total += items[i].Price
		}

		if err := tx.Create(&items).Error; err != nil { // Create order's items
			return err
		}

		if err := tx.Model(&order).Update("price", total).Error; err != nil { // Update order items
			return err
		}

		return nil
	})
}

func ListOrders(db *gorm.DB) ([]models.Order, error) {
	var orders []models.Order
	err := db.Find(&orders).Error
	return orders, err
}

func UpdateOrder(db *gorm.DB, orderID uint, newPrice uint) error {
	return db.Model(&models.Order{}).Where("id = ?", orderID).Update("price", newPrice).Error
}

func DeleteOrder(db *gorm.DB, orderID uint) error {
	return db.Delete(&models.Order{}, orderID).Error
}

// Subquery example
func GetExpensiveOrders(db *gorm.DB) ([]models.Order, error) {
	var orders []models.Order

	/* This is equivalent as this query, with a subquery inside:
	SELECT * from orders
	WHERE price > (SELECT AVG(price) FROM orders)
	LIMIT 20;
	*/

	err := db.Where("price > (?)", db.Table("orders").Select("AVG(price)")).
		Limit(20).
		Find(&orders).
		Error

	return orders, err
}

func GetOrdersWithItems(db *gorm.DB) ([]models.Order, error) {
	var orders []models.Order
	err := db.Preload("Items").Find(&orders).Error
	return orders, err
}
