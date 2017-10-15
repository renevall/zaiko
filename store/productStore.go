package store

import (
	"errors"

	"github.com/renevall/zaiko/domain"
)

//CreateProduct creates a Product Object inside the db store
func (db *DB) CreateProduct(p *domain.Product) {
	db.Create(p)
}

//ReadProductByID creates a Product Object inside the db store
func (db *DB) ReadProductByID(id int) domain.Product {
	var product domain.Product
	db.First(&product, id)
	return product
}

//ReadProductByName creates a Product Object inside the db store
func (db *DB) ReadProductByName(name string) domain.Product {
	var product domain.Product
	db.First(&product, "Name = ?", name)
	return product
}

//UpdateProduct creates a Product Object inside the db store
func (db *DB) UpdateProduct(p domain.Product) domain.Product {
	db.Save(p)
	return p
}

//DeleteProduct creates a Product Object inside the db store
func (db *DB) DeleteProduct(id int) error {

	var p domain.Product
	db.First(&p, "id = ?", id)
	if p.ID <= 0 {
		return errors.New("Object not found")
	}

	err := db.Delete(p).Error
	if err != nil {
		return err
	}

	return nil
}
