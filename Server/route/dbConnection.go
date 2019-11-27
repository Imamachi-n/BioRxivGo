package route

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connectDB() (*gorm.DB, error) {
	// Connect with postgreSQL Database
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=biorxiv user=postgres password=postgres sslmode=disable")

	return db, err
}

func closeDB(db *gorm.DB) {
	db.Close()
}
