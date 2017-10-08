package store

import (
	"github.com/jinzhu/gorm"
	"github.com/renevall/zaiko"
)

type DB struct {
	*gorm.DB
}

//NewDB initializes the database
func NewDB() *DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&zaiko.Product{}, &zaiko.Stock{}, &zaiko.Transaction{})

	return &DB{db}
}
