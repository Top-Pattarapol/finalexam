package database

import (
	"database/sql"

	"github.com/Top-Pattarapol/finalexam/model"
)

func CreateCustomerTable() error {
	db := connect()
	defer db.Close()
	return baseExec(db, `CREATE TABLE IF NOT EXISTS customers( id SERIAL PRIMARY KEY, name TEXT, email TEXT, status TEXT );`)
}

func GetCustomers() (*sql.Rows, error) {
	db := connect()
	defer db.Close()
	return baseQuery(db, `Select id, name, email, status FROM customers ORDER BY id ASC`)
}

func GetCustomerById(id int) (*sql.Row, error) {
	db := connect()
	defer db.Close()
	return baseQueryRow(db, `Select id, name, email, status FROM customers WHERE id=$1 ORDER BY id ASC`, id)
}

func PostCustomers(c *model.Customer) (*sql.Row, error) {
	db := connect()
	defer db.Close()
	return baseQueryRow(db, `INSERT INTO customers (name, email, status) VALUES ($1, $2, $3) RETURNING id`, c.Id, c.Name, c.Email)
}

func DeleteCustomerById(id int) error {
	db := connect()
	defer db.Close()
	return baseExec(db, `DELETE FROM customers WHERE id=$1`, id)
}

func UpdateCustomer(id int, name string, email string, status string) error {
	db := connect()
	defer db.Close()
	return baseExec(db, `UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1`, id, name, email, status)
}
