package services

import (
	"github.com/HenryNg101/golang-examples/sql_db/pkg/models"
	"gorm.io/gorm"
)

type UserSpending struct {
	UserID uint
	Total  uint
}

func GetTopSpenders(db *gorm.DB) ([]UserSpending, error) {
	var result []UserSpending

	err := db.Raw(`
		SELECT user_id, SUM(price) as total
		FROM orders
		GROUP BY user_id
		ORDER BY total DESC
		LIMIT 10
	`).Scan(&result).Error

	return result, err
}

func GetUserOrders(db *gorm.DB, userID uint) (models.User, error) {
	var user models.User

	err := db.Preload("Orders.Items").
		First(&user, userID).Error

	return user, err
}

func CreateUser(db *gorm.DB, name string) (models.User, error) {
	user := models.User{Name: name}
	err := db.Create(&user).Error
	return user, err
}

func ListUsers(db *gorm.DB) (int64, error) {
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
