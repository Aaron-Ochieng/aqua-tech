package main

import (
	"encoding/json"
	"fmt"
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
		// Check if the method is POST
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed) // 405 Method Not Allowed
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Only POST method is allowed",
			})
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		defer r.Body.Close()

		var data Data

		if err = json.Unmarshal(body, &data); err != nil {
			fmt.Println(err)
		}

		fmt.Println(data)
		w.WriteHeader(http.StatusOK)
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
