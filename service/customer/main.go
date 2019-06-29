package customer

import (
	"net/http"

	"github.com/Top-Pattarapol/finalexam/database"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	database *database.Handler
}

func (h *Handler) Init() {
	database := &database.Handler{}
	h.database = database
	h.database.Open()
	h.CreateCustomerTable()
}

func (h *Handler) Close() {
	h.database.Close()
}

func (h *Handler) AuthMiddlewere(c *gin.Context) {
	if token := c.GetHeader("Authorization"); token != "token2019" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	c.Next()
}
