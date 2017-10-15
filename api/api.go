package api

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/renevall/zaiko/store"
)

//NewAPI starts the router
func NewAPI(db *store.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	setRoutes(db, router)

	return router
}

func setRoutes(db *store.DB, r *gin.Engine) {
	r.GET("/product/:id", NotImplemented)
	r.POST("/product", CreateProduct(db))
	r.PUT("/product/:id", NotImplemented)
	r.PATCH("/product/:id", NotImplemented)
	r.DELETE("/product/:id", DeleteProduct(db))

	r.GET("/trans/:id", NotImplemented)
	r.GET("/trans", NotImplemented)
	r.POST("/trans", NotImplemented)
	r.GET("/product/:id/trans", NotImplemented)
}

//NotImplemented is a test dummy for handler
func NotImplemented(c *gin.Context) {
	content := gin.H{"Response": "Not Implemented"}
	c.JSON(http.StatusOK, content)
}