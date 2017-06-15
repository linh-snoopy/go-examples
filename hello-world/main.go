package main

import (
	"fmt"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Get("/", handle)
	r.Get("/:id", handleIndex)
	http.ListenAndServe(":8080", r)
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	index := chi.URLParam(r, "id")
	ctx := r.Context()
	fmt.Println(ctx)
	key, ok := ctx.Value("id").(string)
	if ok {
		fmt.Print("AAAAAAAAAAAAAAAAAAAAA: ")
		fmt.Println(key)
	} else {
		fmt.Print("BBBBBBBBBBBBBBB: ")
		fmt.Println(ok)
		fmt.Println(key)
	}
	w.Write([]byte(fmt.Sprintf("Hello, %s", index)))
}
