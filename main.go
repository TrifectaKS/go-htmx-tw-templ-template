package main

import (
	"go-htmx-tw-templ-template/templates" // replace with your actual module name
	"net/http"
)

func main() {
	// Serve static files (optional)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route: full page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// For now, serve a plain HTML page or integrate templ as needed
		http.ServeFile(w, r, "templates/index.html")
	})

	// Route: HTMX partial using templ
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		err := templates.Hello("hello world").Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}
