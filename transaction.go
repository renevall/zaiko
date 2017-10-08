package zaiko

import (
	"github.com/jinzhu/gorm"
)

// Transaction model definition
type Transaction struct {
	gorm.Model
	Product Product
	Amount  int
	Action  string
}
