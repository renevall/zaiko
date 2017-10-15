package store

import (
	"github.com/jinzhu/gorm"
	"github.com/renevall/zaiko/domain"
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

	// Migrate the schema
	db.AutoMigrate(&domain.Product{}, &domain.Stock{}, &domain.Trans{})

	return &DB{db}
}
