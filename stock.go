package zaiko

import (
	"github.com/jinzhu/gorm"
)

//Stock model definition
type Stock struct {
	gorm.Model
	Product Product
	Amount  int
}
