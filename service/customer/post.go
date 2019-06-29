package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Post(c *gin.Context) {

	t := &Customer{}

	if err := c.BindJSON(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	var id int
	row, err := h.Db.PostCustomers(t.Id, t.Name, t.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
	}
	err = row.Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	t.Id = id
	c.JSON(http.StatusCreated, t)

}
