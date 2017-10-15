package store

import "github.com/renevall/zaiko/domain"

//CreateTrans creates a Trans Object inside the db store
func (db *DB) CreateTrans(t *domain.Trans) {
	db.Create(t)
}

//ReadTransByID creates a Trans Object inside the db store
func (db *DB) ReadTransByID(id int) domain.Trans {
	var Trans domain.Trans
	db.First(&Trans, id)
	return Trans
}

//ReadTransByProduct creates a Trans Object inside the db store
func (db *DB) ReadTransByProduct(id int) []domain.Trans {
	var t []domain.Trans
	db.Where("productID = ?", id).Find(&t)
	return t
}

//UpdateTrans creates a Trans Object inside the db store
func (db *DB) UpdateTrans(t domain.Trans) domain.Trans {
	db.Save(t)
	return t
}
