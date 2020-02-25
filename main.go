package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	log.Println("Starting server...")
	router := chi.NewRouter()

	static := "./static"

	log.Println("Creating file server...")
	fs := http.FileServer(http.Dir(static))

	log.Println("Mapping routes...")
	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(static + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	log.Println("Starting server...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	port = ":" + port
	log.Fatal(http.ListenAndServe(port, router))
}
