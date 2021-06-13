package main

import (
	"log"
	"net/http"
	"news_portal/internal/app/news"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", news.Index)
	r.HandleFunc("/politics", news.Politics)
	r.HandleFunc("/economy", news.Economy)
	r.HandleFunc("/science", news.Science)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:10000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

