package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Handler struct {
	Db *sql.DB
}

func (h *Handler) Open() {
	database, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	h.Db = database
}

func (h *Handler) Close() {
	h.Db.Close()
}

func (h *Handler) BaseExec(query string, args ...interface{}) error {
	stmt, err := h.Db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) BaseQuery(query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := h.Db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func (h *Handler) BaseQueryRow(query string, args ...interface{}) (*sql.Row, error) {
	stmt, err := h.Db.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(args...)
	return row, nil
}
