package main

import (
	"fmt"
	"net/http"

	"github.com/Top-Pattarapol/finalexam/database"
	"github.com/Top-Pattarapol/finalexam/service/customer"

	"github.com/gin-gonic/gin"
)

func main() {
	database.CreateCustomerTable()
	r := serRoute()
	r.Run(getPort())
}

func serRoute() *gin.Engine {
	r := gin.Default()
	r.Use(authMiddlewere)
	r.GET("/customers", customer.Get)
	r.GET("/customers/:id", customer.GetById)
	r.POST("/customers", customer.Post)
	r.DELETE("/customers/:id", customer.DeleteById)
	r.PUT("/customers/:id", customer.Update)
	return r
}

func authMiddlewere(c *gin.Context) {
	fmt.Println("begining middlewere")
	if token := c.GetHeader("Authorization"); token != "token2019" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	c.Next()
	fmt.Println("after end middlewere")
}

func getPort() string {
	return ":2019"
}
