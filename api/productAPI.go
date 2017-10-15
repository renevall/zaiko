package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renevall/domain"
)

type idForm struct {
	id int
}

//CreateProduct creates a product object inside the db
func CreateProduct(db domain.ProductStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		//bind json request to object (worth creating a local object?)
		var p domain.Product
		if c.BindJSON(&p) != nil {
			fmt.Printf("%+v\n", p)
			fmt.Println("no binding")

			c.JSON(http.StatusUnauthorized, gin.H{"status": "product not found"})
			return
		}

		if p.Name != "" {
			ps := domain.NewProduct(db)
			ps.Store.CreateProduct(&p)
			c.JSON(http.StatusOK, gin.H{"status": "product found", "data": p})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"status": "product not found"})
		return

	}
}

//DeleteProduct creates a product from the db
func DeleteProduct(db domain.ProductStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			fmt.Println("no binding")

			c.JSON(http.StatusBadRequest, gin.H{"status": "product not found"})
			return
		}

		ps := domain.NewProduct(db)
		pID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
			return

		}
		err = ps.Store.DeleteProduct(pID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "product not deleted"})
			return

		}
		c.JSON(http.StatusOK, gin.H{"status": "product deleted"})

	}
}
