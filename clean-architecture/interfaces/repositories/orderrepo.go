package repositories

import (
	"github.com/linh-snoopy/go-examples/clean-architecture/entities"
	"fmt"
)

type DbOrderRepo DbRepo

func NewDbOrderRepo(dbHandlers map[string]DbHandler) *DbOrderRepo {
    dbOrderRepo := new(DbOrderRepo)
    dbOrderRepo.dbHandlers = dbHandlers
    dbOrderRepo.dbHandler = dbHandlers["DbOrderRepo"]
    return dbOrderRepo
}

func (repo *DbOrderRepo) Store(order entities.Order) {
    repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO orders (id, customer_id)
                                        VALUES ('%d', '%v')`,
                                        order.Id, order.Customer.Id))
    for _, item := range order.Items {
        repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO items2orders (item_id, order_id)
                                            VALUES ('%d', '%d')`,
                                            item.Id, order.Id))
    }
}

func (repo *DbOrderRepo) FindById(id int) entities.Order {
    row := repo.dbHandler.Query(fmt.Sprintf(`SELECT customer_id FROM orders
                                             WHERE id = '%d' LIMIT 1`,
                                             id))
    var customerId int
    row.Next()
    row.Scan(&customerId)
    customerRepo := NewDbCustomerRepo(repo.dbHandlers)
    order := entities.Order{Id: id, Customer: customerRepo.FindById(customerId)}
    var itemId int
    itemRepo := NewDbItemRepo(repo.dbHandlers)
    row = repo.dbHandler.Query(fmt.Sprintf(`SELECT item_id FROM items2orders
                                            WHERE order_id = '%d'`,
                                            order.Id))
    for row.Next() {
        row.Scan(&itemId)
        order.Add(itemRepo.FindById(itemId))
    }
    return order
}