package repositories

type DbHandler interface {
	Execute(statement string)
	Query(statement string) Rows
	QueryRow(statement string) Row
}

type Rows interface {
	Scan(dest ...interface{})
	Next() bool
}

type Row interface {
	Scan(dest ...interface{})
}

type DbRepo struct {
	dbHandlers map[string]DbHandler
	dbHandler  DbHandler
}
