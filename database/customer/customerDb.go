package customerDb

import (
	"database/sql"
	"log"

	"github.com/Top-Pattarapol/finalexam/database"
)

type Handler struct {
	Db *sql.DB
}

func (h *Handler) CreateCustomerTable() {
	err := database.BaseExec(h.Db, `CREATE TABLE IF NOT EXISTS customers( id SERIAL PRIMARY KEY, name TEXT, email TEXT, status TEXT );`)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (h *Handler) GetCustomers() (*sql.Rows, error) {
	return database.BaseQuery(h.Db, `Select id, name, email, status FROM customers ORDER BY id ASC`)
}

func (h *Handler) GetCustomerById(id int) (*sql.Row, error) {
	return database.BaseQueryRow(h.Db, `Select id, name, email, status FROM customers WHERE id=$1 ORDER BY id ASC`, id)
}

func (h *Handler) PostCustomers(id int, name string, email string) (*sql.Row, error) {
	return database.BaseQueryRow(h.Db, `INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id`, id, name, email)
}

func (h *Handler) DeleteCustomerById(id int) error {
	return database.BaseExec(h.Db, `DELETE FROM customers WHERE id=$1`, id)
}

func (h *Handler) UpdateCustomer(id int, name string, email string, status string) error {
	return database.BaseExec(h.Db, `UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1`, id, name, email, status)
}
