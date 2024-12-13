package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

var (
	t   *template.Template
	err error
)

func init() {
	t, err = t.ParseGlob("templates/*.html")
	if err != nil {
		return
	}
}
func RenderTemplates(w http.ResponseWriter, page string) {
	t.ExecuteTemplate(w, page, nil)
}

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplates(w, "dashboard.html")
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		// Check if the method is POST
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed) // 405 Method Not Allowed
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}

		// Example response payload
		response := map[string]interface{}{
			"message": "Data received successfully",
			"status":  "success",
			"data":    "Your payload goes here",
		}

		// Set response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Encode the response to JSON and send it
		json.NewEncoder(w).Encode(response)
	})
	return mux
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: Router(),
	}

	log.Println("Server running on http://localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}
