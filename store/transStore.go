package store

import (
	"github.com/pkg/errors"

	"github.com/renevall/zaiko/domain"
)

//CreateTrans creates a Trans Object inside the db store
func (db *DB) CreateTrans(t *domain.Trans) error {
	err := db.Create(t).Error
	if err != nil {
		return errors.Wrap(err, "Trying to save transaction")
	}
	return nil
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
