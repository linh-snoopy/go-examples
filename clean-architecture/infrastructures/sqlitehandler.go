package infrastructures

import (
	"database/sql"
	"fmt"
	"github.com/linh-snoopy/go-examples/clean-architecture/interfaces/repositories"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteHandler struct {
	Conn *sql.DB
}

func (handler *SqliteHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *SqliteHandler) Query(statement string) repositories.Rows {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(SqliteRows)
	}
	row := new(SqliteRows)
	row.Rows = rows
	return row
}

func (handler *SqliteHandler) QueryRow(statement string) repositories.Row {
	row := handler.Conn.QueryRow(statement)
	r := new(SqliteRow)
	r.Row = row
	return r
}

// type sql.Row is the result of calling QueryRow to select single row.
// type sql.Rows is the result of a query. Its cursor start before the first row of the result.
type SqliteRows struct {
	Rows *sql.Rows
}

type SqliteRow struct {
	Row *sql.Row
}

func (r SqliteRows) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r SqliteRows) Next() bool {
	return r.Rows.Next()
}

func (r SqliteRow) Scan(dest ...interface{}) {
	r.Row.Scan(dest...)
}

func NewSqliteHandler(dbFileName string) *SqliteHandler {
	conn, _ := sql.Open("sqlite3", dbFileName)
	sqliteHandler := &SqliteHandler{Conn: conn}
	return sqliteHandler
}
