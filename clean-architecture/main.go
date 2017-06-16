package main

import (
	"github.com/linh-snoopy/go-examples/clean-architecture/interfaces"
	"github.com/linh-snoopy/go-examples/clean-architecture/interfaces/repositories"
	//	"github.com/linh-snoopy/go-examples/clean-architecture/entities"
	"github.com/linh-snoopy/go-examples/clean-architecture/infrastructures"
	"github.com/linh-snoopy/go-examples/clean-architecture/usecases"
)

var DB = "postgres://postgres:postgres@localhost/tse_clean?sslmode=disable"

func main() {
	dbHandler := infrastructures.NewPostgresHandler(DB)

	handlers := make(map[string]interfaces.DBHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbCustomerRepo"] = dbHandler
	handlers["DbItemRepo"] = dbHandler
	handlers["DbOrderRepo"] = dbHandler

	orderInteractor := new(usecases.OrderInteractor)
	orderInteractor.UserRepository = interfaces.NewDbUserRepo(handlers)
	orderInteractor.ItemRepository = interfaces.NewDbItemRepo(handlers)
	orderInteractor.OrderRepository = interfaces.NewDbOrderRepo(handlers)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.OrderInteractor = orderInteractor

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowOrder(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
