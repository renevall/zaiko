package domain

import "github.com/jinzhu/gorm"

//Product model definition
type Product struct {
	gorm.Model
	Name  string       `json:"name"`
	Store ProductStore `gorm:"-" json:"-"` //Interface
	Trans []Trans
}

//ProductStore is the contract interface to the DB
type ProductStore interface {
	CreateProduct(*Product)
	ReadProductByID(int) Product
	ReadProductByName(string) Product
	UpdateProduct(Product) Product
	DeleteProduct(id int) error
}

//NewProduct returns an instance with the db/store dependency met
func NewProduct(ps ProductStore) *Product {
	return &Product{
		Store: ps,
	}
}
