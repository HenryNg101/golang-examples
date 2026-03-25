package main

import (
	"fmt"
	"log"

	"github.com/HenryNg101/golang-examples/sql_db/pkg/db"
	"gorm.io/gorm"
)

func RunExperiments(db *gorm.DB, exps []QueryExperiment) {
	for _, e := range exps {
		// Setup
		for _, sql := range e.SetupSQL {
			if err := db.Exec(sql).Error; err != nil {
				log.Fatalf("[%s] Setup failed: %v", e.Name, err)
			}
		}

		// Benchmark
		var result []map[string]interface{}
		fmt.Printf("=== Running experiment: %s ===\n", e.Name)
		if err := db.Raw("EXPLAIN ANALYZE " + e.QuerySQL).Scan(&result).Error; err != nil {
			log.Fatalf("[%s] Query failed: %v", e.Name, err)
		}

		// Print results
		for _, row := range result {
			// EXPLAIN ANALYZE returns a single column called "QUERY PLAN"
			fmt.Println(row["QUERY PLAN"])
		}

		// Teardown
		for _, sql := range e.TeardownSQL {
			if err := db.Exec(sql).Error; err != nil {
				log.Fatalf("[%s] Teardown failed: %v", e.Name, err)
			}
		}
		fmt.Println()
	}
}

func main() {
	db := db.ConnectDB()
	experiments := []QueryExperiment{
		{
			Name:     "No Index",
			QuerySQL: `SELECT order_id, SUM(price) FROM order_items GROUP BY order_id`,
		},
		{
			Name:        "Index on order_id",
			SetupSQL:    []string{`CREATE INDEX idx_order_items_orderid ON order_items(order_id)`},
			QuerySQL:    `SELECT order_id, SUM(price) FROM order_items GROUP BY order_id`,
			TeardownSQL: []string{`DROP INDEX idx_order_items_orderid`},
		},
		{
			Name:        "Index on price",
			SetupSQL:    []string{`CREATE INDEX idx_order_items_price ON order_items(price)`},
			QuerySQL:    `SELECT order_id, SUM(price) FROM order_items GROUP BY order_id`,
			TeardownSQL: []string{`DROP INDEX idx_order_items_price`},
		},
		{
			Name: "Separate indexes order_id + price",
			SetupSQL: []string{
				`CREATE INDEX idx_order_items_orderid ON order_items(order_id)`,
				`CREATE INDEX idx_order_items_price ON order_items(price)`,
			},
			QuerySQL: `SELECT order_id, SUM(price) FROM order_items GROUP BY order_id`,
			TeardownSQL: []string{
				`DROP INDEX idx_order_items_orderid`,
				`DROP INDEX idx_order_items_price`,
			},
		},
		{
			Name:        "Combined index order_id, price",
			SetupSQL:    []string{`CREATE INDEX idx_order_items_orderid_price ON order_items(order_id, price)`},
			QuerySQL:    `SELECT order_id, SUM(price) FROM order_items GROUP BY order_id`,
			TeardownSQL: []string{`DROP INDEX idx_order_items_orderid_price`},
		},
		{
			Name:        "Combined index price, order_id",
			SetupSQL:    []string{`CREATE INDEX idx_order_items_price_orderid ON order_items(price, order_id)`},
			QuerySQL:    `SELECT order_id, SUM(price) FROM order_items GROUP BY order_id`,
			TeardownSQL: []string{`DROP INDEX idx_order_items_price_orderid`},
		},
	}
	RunExperiments(db, experiments)
}
