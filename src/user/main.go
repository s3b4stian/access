package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	host := "0.0.0.0:3000"

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         host,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("[+] server listening at " + host)

	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"status\":\"ok\"}")
}
