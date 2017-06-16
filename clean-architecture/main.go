package main

import (
	"github.com/linh-snoopy/go-examples/clean-architecture/interfaces"
	"github.com/linh-snoopy/go-examples/clean-architecture/interfaces/repositories"
	//	"github.com/linh-snoopy/go-examples/clean-architecture/entities"
	"github.com/linh-snoopy/go-examples/clean-architecture/infrastructures"
	"github.com/linh-snoopy/go-examples/clean-architecture/usecases"
	"net/http"
)

var DB = "postgres://postgres:postgres@localhost/test_clean?sslmode=disable"

func main() {
	dbHandler, err := infrastructures.NewPostgresHandler(DB)
	if err!= nil {
		panic(err)
	}

	handlers := make(map[string]repositories.DbHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbCustomerRepo"] = dbHandler
	handlers["DbItemRepo"] = dbHandler
	handlers["DbOrderRepo"] = dbHandler

	orderInteractor := new(usecases.OrderInteractor)
	orderInteractor.UserRepository = repositories.NewDbUserRepo(handlers)
	orderInteractor.ItemRepository = repositories.NewDbItemRepo(handlers)
	orderInteractor.OrderRepository = repositories.NewDbOrderRepo(handlers)

	webserviceHandler := interfaces.WebServiceHandler{}
	webserviceHandler.OrderInteractor = orderInteractor

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowOrder(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
