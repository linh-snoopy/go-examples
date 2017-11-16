package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("My API", func() {
	Title("Users Management")
	Scheme("http")
	Host("localhost:8080")
	Consumes("application/json")
	Produces("application/json")
})

var User = Type("user", func() {
	Description("")
	Attribute("email", String, "User's email address")
	Attribute("name", String, "User's name")
	Attribute("phone", String, "Phone's number")
	Required("email", "name")
})

var UserMedia = MediaType("vnd.my.user", func() {
	Reference(User)
	Attributes(func() {
		Attribute("email")
		Attribute("name")
		Attribute("phone")
	})
	View("default", func() {
		Attribute("email")
		Attribute("name")
		Attribute("phone")
	})
})

var Results = MediaType("vnd.my.result", func () {
	Description("The result of an operation")
	Attributes(func() {
		Attribute("value", Integer, "Results value")
	})
	View("extended", func() {
		Attribute("value")
	})
	View("default", func() {
		Attribute("value")
	})
})

var _ = Resource("Users", func() {
	BasePath("/users")
	Action("add222", func() {
		Description("Register a new user")
		Routing(POST("/add223344"))
		Payload(User)
		Response(Created)
	})

	Action("list", func() {
		Description("List all users")
		Routing(GET("/list"))
		Response(OK, CollectionOf(UserMedia))
	})
	
	Action("detail", func() {
		Description("Get detail of user")
		Routing(GET("/detail/:id"))
		Params(func() {
			Param("id", String, "user id")
		})
		Response(OK, UserMedia)
	})
})

var _ = Resource("Operands", func () {
	BasePath("/results")
	Action("sum", func () {
		Description("Sum")
		Routing(GET("/sum/:left/:right"))
		Params(func() {
			Param("left", Integer, "Left operand")
			Param("right", Integer, "Right operand")
		})
		Response(OK)
	})
})