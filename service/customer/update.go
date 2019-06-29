package customer

import (
	"net/http"

	"github.com/Top-Pattarapol/finalexam/utility"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {

	id, err := utility.ParamToInt(c, "id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	t := &Customer{}

	if err := c.BindJSON(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	t.Id = id

	err = h.UpdateCustomer(id, t.Name, t.Email, t.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, t)
}
