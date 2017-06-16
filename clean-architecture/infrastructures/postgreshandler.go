package infrastructures

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/linh-snoopy/go-examples/clean-architecture/interfaces/repositories"
)

type PostgresHandler struct {
	Conn *sql.DB
}

func (handler *PostgresHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *PostgresHandler) Query(statement string) repositories.Rows {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(PostgresRows)
	}
	row := new(PostgresRows)
	row.Rows = rows
	return row
}

func (handler *PostgresHandler) QueryRow(statement string) repositories.Row {
	row := handler.Conn.QueryRow(statement)
	sRow := new(PostgresRow)
	sRow.Row = row
	return sRow
}

// type sql.Row is the result of calling QueryRow to select single row.
// type sql.Rows is the result of a query. Its cursor start before the first row of the result.
type PostgresRows struct {
	Rows *sql.Rows
}

type PostgresRow struct {
	Row *sql.Row
}

func (r PostgresRows) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r PostgresRows) Next() bool {
	return r.Rows.Next()
}

func (r PostgresRow) Scan(dest ...interface{}) {
	r.Row.Scan(dest...)
}

func NewPostgresHandler(dbFileName string) (*PostgresHandler, error) {
	conn, err := sql.Open("postgres", dbFileName)
	if err != nil {
		return nil, err
	}
	return &PostgresHandler{Conn: conn}, nil
}
