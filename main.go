package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	static := "./static"
	fs := http.FileServer(http.Dir(static))
	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(static + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	log.Fatal(http.ListenAndServe(":3000", router))
}
