package customer

import (
	"net/http"

	"github.com/Top-Pattarapol/finalexam/database"
	"github.com/Top-Pattarapol/finalexam/utility"
	"github.com/gin-gonic/gin"
)

func DeleteById(c *gin.Context) {

	id, err := utility.ParamToInt(c, "id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	err = database.DeleteCustomerById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted",
	})
}
