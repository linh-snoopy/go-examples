package repositories

import (
	"github.com/linh-snoopy/go-examples/clean-architecture/entities"
	//	"github.com/linh-snoopy/go-examples/clean-architecture/usecases"
	"fmt"
)

type DbCustomerRepo DbRepo

func NewDbCustomerRepo(dbHandlers map[string]DbHandler) *DbCustomerRepo {
	dbCustomerRepo := new(DbCustomerRepo)
	dbCustomerRepo.dbHandlers = dbHandlers
	dbCustomerRepo.dbHandler = dbHandlers["DbCustomerRepo"]
	return dbCustomerRepo
}

func (repo *DbCustomerRepo) Store(customer entities.Customer) {
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO customers (id, name)
                                        VALUES ('%d', '%v')`,
		customer.Id, customer.Name))
}

func (repo *DbCustomerRepo) FindById(id int) entities.Customer {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT name FROM customers
                                             WHERE id = '%d' LIMIT 1`,
		id))
	var name string
	row.Next()
	row.Scan(&name)
	return entities.Customer{Id: id, Name: name}
}
