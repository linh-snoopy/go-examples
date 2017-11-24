package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the api definition used by goa to generate the api
var _ = API("congo", func() {
	Title("Congo")
	Description("Congo")
	Contact(func() {
		Name("congo")
		Email("congo")
		URL("https://congo.gopheracademy.com")
	})

	Docs(func() {
		Description("Getting Started Guide")
		URL("https://congodocs.gopheracademy.com")
	})
	Host("congo.gopheracademy.com")
	Scheme("https")
	BasePath("/api")

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})

	Origin("http://localhost:5000", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		MaxAge(600)
		Credentials()
	})

	JWTSecurity("jwt", func() {
		Header("Authorization")
		// TokenURL("<a href='https://congo.gopheracademy.com/api/auth/token'>https://congo.gopheracademy.com/api/auth/token</a>")
	})
	BasicAuthSecurity("password", func() {
		Description("Use your own password!")
	})
})