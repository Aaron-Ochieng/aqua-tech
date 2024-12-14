package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

var (
	t   *template.Template
	err error
)

type Data struct {
	Temp           float64
	Humidity       float64
	UltraSonicData float64
}

func init() {
	t, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
}

func RenderTemplates(w http.ResponseWriter, page string, data interface{}) {
	if err := t.ExecuteTemplate(w, page, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Printf("Error rendering template %s: %v", page, err)
	}
}

var latestData interface{}

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplates(w, "dashboard.html", nil)
	})
	mux.HandleFunc("/farms", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplates(w, "farms.html")
	})
	mux.HandleFunc("/market", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplates(w, "marketplace.html")
	})
	mux.HandleFunc("/community", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplates(w, "community.html")
	})
	mux.HandleFunc("/reports", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplates(w, "reports.html")
	})
	mux.HandleFunc("/education", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplates(w, "education.html")
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(latestData)
			return
		}
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}


		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			log.Printf("Error reading request body: %v", err)
			return
		}
		defer r.Body.Close()

		var data Data
		if err := json.Unmarshal(body, &data); err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			log.Printf("Error unmarshalling JSON: %v", err)
			return
		}

		latestData = data
		log.Printf("Received data: %+v", data)
		w.WriteHeader(http.StatusOK)
		// RenderTemplates(w, "dashboard.html", data)
	})
	return mux
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: Router(),
	}

	log.Printf("Server running on http://localhost%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
