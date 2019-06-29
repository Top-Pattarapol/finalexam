package customer

import (
	"database/sql"
	"log"
)

func (h *Handler) CreateCustomerTable() {
	err := h.database.BaseExec(`CREATE TABLE IF NOT EXISTS customers( id SERIAL PRIMARY KEY, name TEXT, email TEXT, status TEXT );`)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (h *Handler) GetCustomers() (*sql.Rows, error) {
	return h.database.BaseQuery(`Select id, name, email, status FROM customers ORDER BY id ASC`)
}

func (h *Handler) GetCustomerById(id int) (*sql.Row, error) {
	return h.database.BaseQueryRow(`Select id, name, email, status FROM customers WHERE id=$1 ORDER BY id ASC`, id)
}

func (h *Handler) PostCustomers(id int, name string, email string) (*sql.Row, error) {
	return h.database.BaseQueryRow(`INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id`, id, name, email)
}

func (h *Handler) DeleteCustomerById(id int) error {
	return h.database.BaseExec(`DELETE FROM customers WHERE id=$1`, id)
}

func (h *Handler) UpdateCustomer(id int, name string, email string, status string) error {
	return h.database.BaseExec(`UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1`, id, name, email, status)
}
