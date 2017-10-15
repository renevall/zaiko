package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renevall/zaiko"
)

func CreateTrans(db zaiko.TransStore) gin.HandleFunc {
	return func(c *gin.Context) {
		var t zaiko.Trans
		if c.BindJSON(&t) != nil {
			fmt.Printf("%+v\n", p)
			fmt.Println("no binding")

			c.JSON(http.StatusUnauthorized, gin.H{"status": "product not found"})
			return
		}

		if t.ProductID <= 0 { //no productid
			c.JSON(http.StatusBadRequest, gin.H{"status": "no product ID"})
			return
		}

		if t.Amount <=0 || t.Action = "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "no action done"})
			return
		}
	}
}
