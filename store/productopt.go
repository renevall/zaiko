package store

import "github.com/renevall/zaiko"

//CreateProduct creates a Product Object inside the db store
func (db *DB) CreateProduct(p zaiko.Product) zaiko.Product {
	db.Create(p)
	return p
}

//ReadProductByID creates a Product Object inside the db store
func (db *DB) ReadProductByID(id int) zaiko.Product {
	var product zaiko.Product
	db.First(&product, id)
	return product
}

//ReadProductByName creates a Product Object inside the db store
func (db *DB) ReadProductByName(name string) zaiko.Product {
	var product zaiko.Product
	db.First(&product, "Name = ?", name)
	return product
}

//UpdateProduct creates a Product Object inside the db store
func (db *DB) UpdateProduct() {

}

//DeleteProduct creates a Product Object inside the db store
func (db *DB) DeleteProduct() {

}
