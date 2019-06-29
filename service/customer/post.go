package customer

import (
	"net/http"

	"github.com/Top-Pattarapol/finalexam/database"
	"github.com/Top-Pattarapol/finalexam/model"
	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {

	t := &model.Customer{}

	if err := c.BindJSON(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	var id int
	row, err := database.PostCustomers(t)
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
