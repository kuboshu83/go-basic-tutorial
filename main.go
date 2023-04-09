package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go!!")
}

func GoodByeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Good Bye!!")
}

func AddHandlerLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("start handler")
		next.ServeHTTP(w, r)
		log.Printf("end handler")
	})
}

func main() {
	mux := chi.NewRouter()
	mux.Route("/hello", func(r chi.Router) {
		r.Use(AddHandlerLog)
		r.Get("/", HelloHandler)
	})
	mux.Route("/goodbye", func(r chi.Router) {
		r.Get("/", GoodByeHandler)
	})
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}
