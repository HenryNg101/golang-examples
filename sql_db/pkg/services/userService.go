package services

import (
	"github.com/HenryNg101/golang-examples/sql_db/pkg/models"
	"gorm.io/gorm"
)

type UserSpending struct {
	UserID uint
	Total  uint
}

func GetTopSpenders(db *gorm.DB, count int) ([]UserSpending, error) {
	var result []UserSpending

	/* This query is equal to:
	SELECT user_id, SUM(price) as total
	FROM orders
	GROUP BY user_id
	ORDER BY total asc, user_id asc
	LIMIT ...
	*/

	err := db.Model(&models.Order{}).
		Select("user_id, SUM(price) as total").
		Group("user_id").
		Order("total DESC, user_id ASC").
		Limit(count).
		Find(&result).
		Error

	return result, err
}

func GetUserOrders(db *gorm.DB, userID uint) (models.User, error) {
	var user models.User

	// For each users. get all of it's orders, and for each order, preload all of it's items. Basically, it's nested preloading
	/* So firstly, it will check that there's First(&user, userID), and it will do first query:
	SELECT * FROM users WHERE id = ? ORDER BY id LIMIT 1;

	After that, it loads all the orders:
	SELECT * FROM "orders" WHERE "orders"."user_id" = ?

	Then finally, loads all order items:
	SELECT * FROM "order_items" WHERE "order_items"."order_id" IN (?, ?, ...)

	*/
	err := db.Preload("Orders.OrderItems").
		First(&user, userID).Error

	return user, err
}

func CreateUser(db *gorm.DB, name string) (models.User, error) {
	user := models.User{Name: name}
	err := db.Create(&user).Error
	return user, err
}

func CountUsers(db *gorm.DB) (int64, error) {
	var listCount int64
	err := db.Model(&models.User{}).Count(&listCount).Error
	return listCount, err
}

func UpdateUser(db *gorm.DB, userID uint, newName string) error {
	return db.Model(&models.User{}).Where("id = ?", userID).Update("name", newName).Error
}

func DeleteUser(db *gorm.DB, userID uint) error {
	return db.Delete(&models.User{}, userID).Error
}
