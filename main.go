package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login")
}

func LoginSuccessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login!!")
}

func NotLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Login!!")
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
	mux.Route("/login", func(r chi.Router) {
		r.Get("/", LoginHandler)
	})
	mux.Route("/loginsuccess", func(r chi.Router) {
		r.Use(AddHandlerLog)
		r.Get("/", LoginSuccessHandler)
	})
	mux.Route("/notlogin", func(r chi.Router) {
		r.Get("/", NotLoginHandler)
	})
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}
