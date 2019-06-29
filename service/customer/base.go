package customer

import (
	"net/http"

	customerDb "github.com/Top-Pattarapol/finalexam/database/customer"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Db *customerDb.Handler
}

func (h *Handler) AuthMiddlewere(c *gin.Context) {
	if token := c.GetHeader("Authorization"); token != "token2019" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	c.Next()
}
