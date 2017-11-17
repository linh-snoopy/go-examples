package controllers

import (
	"github.com/goadesign/goa"
	"github.com/linh-snoopy/go-examples/goatest/gen/app"
)

// UsersController implements the Users resource.
type UsersController struct {
	*goa.Controller
}

// NewUsersController creates a Users controller.
func NewUsersController(service *goa.Service) *UsersController {
	return &UsersController{Controller: service.NewController("UsersController")}
}

// Add222 runs the add222 action.
func (c *UsersController) Add222(ctx *app.Add222UsersContext) error {
	// UsersController_Add222: start_implement

	// Put your logic here

	// UsersController_Add222: end_implement
	return nil
}

// Detail runs the detail action.
func (c *UsersController) Detail(ctx *app.DetailUsersContext) error {
	// UsersController_Detail: start_implement

	// Put your logic here
	// log.Println(ctx.ID)

	// UsersController_Detail: end_implement
	res := &app.MyUser{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *UsersController) List(ctx *app.ListUsersContext) error {
	// UsersController_List: start_implement

	// Put your logic here

	// UsersController_List: end_implement
	res := app.MyUserCollection{}
	return ctx.OK(res)
}
