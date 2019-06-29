package customer

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Top-Pattarapol/finalexam/model"

	"github.com/Top-Pattarapol/finalexam/database"
	"github.com/gin-gonic/gin"
)

func DeleteCustomersById(c *gin.Context) {

	id, err := paramToInt(c, "id")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = database.DeleteCustomerById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted",
	})
}

func PostCustomers(c *gin.Context) {

	t := &model.Customer{}

	if err := c.BindJSON(t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var id int
	row, err := database.PostCustomers(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	err = row.Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	t.Id = id
	c.JSON(http.StatusCreated, t)

}

func GetCustomersById(c *gin.Context) {

	id, err := paramToInt(c, "id")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	row, err := database.GetCustomerById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	t := model.Customer{}
	err = row.Scan(&t.Id, &t.Name, &t.Email, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, t)
}

func GetTodos(c *gin.Context) {

	rows, err := database.GetCustomers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todos := []model.Customer{}

	for rows.Next() {
		t := model.Customer{}
		err := rows.Scan(&t.Id, &t.Name, &t.Email, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, t)
	}

	fmt.Println(todos)
	c.JSON(http.StatusOK, todos)
}

func UpdateCustomer(c *gin.Context) {

	id, err := paramToInt(c, "id")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t := &model.Customer{}

	if err := c.BindJSON(t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t.Id = id

	err = database.UpdateCustomer(id, t.Name, t.Email, t.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}

func paramToInt(c *gin.Context, key string) (int, error) {
	param := c.Param(key)
	value, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return value, nil
}
