package customer

import (
	"net/http"

	"github.com/Top-Pattarapol/finalexam/utility"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetById(c *gin.Context) {

	id, err := utility.ParamToInt(c, "id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	row, err := h.GetCustomerById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	t := Customer{}
	err = row.Scan(&t.Id, &t.Name, &t.Email, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, t)
}
