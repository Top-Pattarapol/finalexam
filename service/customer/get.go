package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(c *gin.Context) {

	rows, err := h.Db.GetCustomers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	customers := []Customer{}

	for rows.Next() {
		t := Customer{}
		err := rows.Scan(&t.Id, &t.Name, &t.Email, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
			return
		}
		customers = append(customers, t)
	}
	c.JSON(http.StatusOK, customers)
}
