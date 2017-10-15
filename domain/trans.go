package domain

import (
	"github.com/jinzhu/gorm"
)

// Trans model definition
type Trans struct {
	gorm.Model
	ProductID int
	Amount    int
	Action    string
}

//TransStore is the contract interface to the DB
type TransStore interface {
	CreateTrans(t *Trans)
	ReadTransByID(id int) Trans
	ReadTransByProduct(id int) []Trans
	UpdateTrans(t Trans) Trans
}
