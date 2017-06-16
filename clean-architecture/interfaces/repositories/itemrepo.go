package repositories

import (
	"fmt"
	"github.com/linh-snoopy/go-examples/clean-architecture/entities"
)

type DbItemRepo DbRepo

func NewDbItemRepo(dbHandlers map[string]DbHandler) *DbItemRepo {
	dbItemRepo := new(DbItemRepo)
	dbItemRepo.dbHandlers = dbHandlers
	dbItemRepo.dbHandler = dbHandlers["DbItemRepo"]
	return dbItemRepo
}

func (repo *DbItemRepo) Store(item entities.Item) {
	available := "no"
	if item.Available {
		available = "yes"
	}
	repo.dbHandler.Execute(fmt.Sprintf(`INSERT INTO items (id, name, value, available)
                                        VALUES ('%d', '%v', '%f', '%v')`,
		item.Id, item.Name, item.Value, available))
}

func (repo *DbItemRepo) FindById(id int) entities.Item {
	row := repo.dbHandler.Query(fmt.Sprintf(`SELECT name, value, available
                                             FROM items WHERE id = '%d' LIMIT 1`,
		id))
	var name string
	var value float64
	var available string
	row.Next()
	row.Scan(&name, &value, &available)
	item := entities.Item{Id: id, Name: name, Value: value}
	item.Available = false
	if available == "yes" {
		item.Available = true
	}
	return item
}
