package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renevall/zaiko"
)

//NotImplemented is a test dummy for handler
func NotImplemented(c *gin.Context) {
	content := gin.H{"Response": "Not Implemented"}
	c.JSON(http.StatusOK, content)
}

//CreateProduct creates a product object inside the db
func CreateProduct(db zaiko.ProductStore) gin.HandlerFunc {

	return func(c *gin.Context) {
		//bind json request to object (worth creating a local object?)
		var product zaiko.Product
		if c.BindJSON(&product) == nil {
			fmt.Println(product)
			if product.Name != "" {
				ps := zaiko.NewProduct(db)
				p := ps.Store.CreateProduct(product)
				c.JSON(http.StatusOK, gin.H{"status": "product found", "data": p})
				return
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "product not found"})
				return
			}
		} else {
			fmt.Println("no binding")

			c.JSON(http.StatusUnauthorized, gin.H{"status": "product not found"})
			return
		}
	}
}
