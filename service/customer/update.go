package customer

import (
	"net/http"

	"github.com/Top-Pattarapol/finalexam/model"
	"github.com/Top-Pattarapol/finalexam/utility"

	"github.com/Top-Pattarapol/finalexam/database"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {

	id, err := utility.ParamToInt(c, "id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	t := &model.Customer{}

	if err := c.BindJSON(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	t.Id = id

	err = database.UpdateCustomer(id, t.Name, t.Email, t.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, t)
}
