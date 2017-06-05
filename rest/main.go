package main 

import (
	//"fmt"
	"flag"
	"net/http"
	"github.com/pressly/chi"
	//"github.com/pressly/chi/docgen"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
)

var routes = flag.Bool("routes", false, "Generate router document")

func main() {
	
	flag.Parse()
	r := chi.NewRouter()
	
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	
	r.Get("/", rootHandle)
	r.Get("/ping", pingHandle)
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})
	
	// RESTy routes for "articles" resource
	r.Mount("/articles", articlesRoutes())
	
	http.ListenAndServe(":8080", r)
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome root!"))
}

func pingHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func articlesRoutes() chi.Router {
	r := chi.NewRouter()
	r.With(paginate).Get("/", ListArticles)
	r.Post("/", CreateArticle)				// POST /articles
	r.Get("/search", SearchArticles)		// GET  /articles/search
	return r
}

func ListArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ListArticles"))
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateArticle"))
}

func SearchArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SearchArticles"))
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}