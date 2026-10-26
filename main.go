//go:generate go tool templ generate

package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"raplab/views"
)

//go:embed static
var staticFiles embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	static, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := views.Home().Render(r.Context(), w); err != nil {
			log.Printf("render homepage: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	addr := ":" + port
	log.Printf("RAP-LAB listening on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
