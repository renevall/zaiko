package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renevall/zaiko/domain"
)

func CreateTrans(db domain.TransStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var t domain.Trans
		if c.BindJSON(&t) != nil {
			fmt.Printf("%+v\n", t)
			fmt.Println("no binding")

			c.JSON(http.StatusUnauthorized, gin.H{"status": "product not found"})
			return
		}

		if t.ProductID <= 0 { //no productid
			c.JSON(http.StatusBadRequest, gin.H{"status": "no product ID"})
			return
		}

		if t.Amount <= 0 || t.Action == "" {
			fmt.Printf("%+v\n", t)
			c.JSON(http.StatusBadRequest, gin.H{"status": "no action done"})
			return
		}

		err := db.CreateTrans(&t)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
			fmt.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Transaction Created", "data": t})
	}
}
