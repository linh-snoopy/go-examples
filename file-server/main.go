package main

import (
	"fmt"
	"github.com/pressly/chi"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))
	})
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "files")
	r.FileServer("/files", http.Dir(filesDir))

	fmt.Println(filesDir)
	http.ListenAndServe(":8080", r)
}
