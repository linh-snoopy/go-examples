package repositories

import (
//	"github.com/linh-snoopy/go-examples/clean-architecture/entities"
	"github.com/linh-snoopy/go-examples/clean-architecture/usecases"
	"fmt"
)

type DbUserRepo DbRepo

func NewDbUserRepo(dbHandlers map[string]DbHandler) *DbUserRepo {
    dbUserRepo := new(DbUserRepo)
    dbUserRepo.dbHandlers = dbHandlers
    dbUserRepo.dbHandler = dbHandlers["DbUserRepo"]
    return dbUserRepo
}

func (repo *DbUserRepo) Store(user usecases.User) {
    isAdmin := "no"
    if user.IsAdmin {
        isAdmin = "yes"
    }
    repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO users (id, customer_id, is_admin)
                                        VALUES ('%d', '%d', '%v')`,
                                        user.Id, user.Customer.Id, isAdmin))
    customerRepo := NewDbCustomerRepo(repo.dbHandlers)
    customerRepo.Store(user.Customer)
}

func (repo *DbUserRepo) FindById(id int) usecases.User {
    row := repo.dbHandler.Query(fmt.Sprintf(`SELECT is_admin, customer_id
                                             FROM users WHERE id = '%d' LIMIT 1`,
                                             id))
    var isAdmin string
    var customerId int
    row.Next()
    row.Scan(&isAdmin, &customerId)
    customerRepo := NewDbCustomerRepo(repo.dbHandlers)
    u := usecases.User{Id: id, Customer: customerRepo.FindById(customerId)}
    u.IsAdmin = false
    if isAdmin == "yes" {
        u.IsAdmin = true
    }
    return u
}