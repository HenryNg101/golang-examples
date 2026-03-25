package services

import (
	"github.com/HenryNg101/golang-examples/sql_db/pkg/models"
	"gorm.io/gorm"
)

func CreateOrderItem(db *gorm.DB, orderID uint, price uint, name string) (models.OrderItem, error) {
	item := models.OrderItem{
		OrderID: orderID,
		Price:   price,
		Name:    name,
	}
	err := db.Create(&item).Error
	return item, err
}

func ListOrderItems(db *gorm.DB) ([]models.OrderItem, error) {
	var items []models.OrderItem
	err := db.Find(&items).Error
	return items, err
}

func UpdateOrderItem(db *gorm.DB, itemID uint, newPrice uint, newName string) error {
	return db.Model(&models.OrderItem{}).
		Where("id = ?", itemID).
		Updates(map[string]interface{}{
			"price": newPrice,
			"name":  newName,
		}).Error
}

func DeleteOrderItem(db *gorm.DB, itemID uint) error {
	return db.Delete(&models.OrderItem{}, itemID).Error
}
