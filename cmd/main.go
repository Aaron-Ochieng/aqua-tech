package main

import (
	"log"
	"net/http"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Home"))
	})
	return mux
}
func main(){
	server := http.Server{
		Addr: ":8080",
		Handler: Router() ,
	}

	log.Println("Server running on http://localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}