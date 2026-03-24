package pkg

import (
	"log"

	"gorm.io/gorm"
)

func BulkCreate(db *gorm.DB) {
	err := db.Exec(`
	INSERT INTO users(name, created_at)
	SELECT
		'user_' || g,
		NOW() - (g || ' seconds')::interval
	FROM generate_series(1,1000000) g`).Error

	if err != nil {
		log.Fatal(err)
	}

	err = db.Exec(`
	INSERT INTO orders(user_id, price, created_at)
	SELECT
		(floor(random()*1000000))::int + 1,
		(floor(random()*500))::int,
		NOW() - (g || ' seconds')::interval
	FROM generate_series(1,5000000) g`).Error

	if err != nil {
		log.Fatal(err)
	}

	err = db.Exec(`
	INSERT INTO order_items(order_id, price, name, created_at)
	SELECT
		(floor(random()*5000000))::int + 1,
		(floor(random()*500))::int,
		'default',
		NOW() - (g || ' seconds')::interval
	FROM generate_series(1,10000000) g`).Error

	if err != nil {
		log.Fatal(err)
	}
}
