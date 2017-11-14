//go:generate goagen bootstrap -d github.com/linh-snoopy/go-examples/goatest/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/linh-snoopy/go-examples/goatest/app"
)

func main() {
	// Create service
	service := goa.New("My API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "Operands" controller
	c := NewOperandsController(service)
	app.MountOperandsController(service, c)
	// Mount "Users" controller
	c2 := NewUsersController(service)
	app.MountUsersController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
