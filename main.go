package main

import (
	"github.com/Top-Pattarapol/finalexam/database"
	"github.com/Top-Pattarapol/finalexam/service/customer"

	"github.com/gin-gonic/gin"
)

func main() {
	db := &database.Handler{}
	db.Db = database.Connect()
	defer db.Db.Close()

	h := &customer.Handler{}
	h.Db = db

	r := serRoute(h)
	r.Run(getPort())
}

func serRoute(h *customer.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(h.AuthMiddlewere)
	r.GET("/customers", h.Get)
	r.GET("/customers/:id", h.GetById)
	r.POST("/customers", h.Post)
	r.DELETE("/customers/:id", h.DeleteById)
	r.PUT("/customers/:id", h.Update)
	return r
}

func getPort() string {
	return ":2019"
}
