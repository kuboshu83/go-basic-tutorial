package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func checkBasicAuth(r *http.Request) bool {
	userName, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	if userName != "kuboshu" || password != "kuboshu" {
		return false
	}
	return true
}

func RedirectNotAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authOk := checkBasicAuth(r)
		if !authOk {
			// PermanentRedirectにするとchromeがキャッシュして、ログイン後も勝手にリダイレクトしてハマった
			http.Redirect(w, r, "http://localhost:8080/login", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	authOk := checkBasicAuth(r)
	if !authOk {
		w.Header().Add("WWW-Authenticate", "Basic")
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "already authorized")
}

func PrivatePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is private zone!!")
}

func main() {
	mux := chi.NewRouter()
	mux.Route("/login", func(r chi.Router) {
		r.Get("/", LoginHandler)
	})
	mux.Route("/private", func(r chi.Router) {
		r.Use(RedirectNotAuthorized)
		r.Get("/", PrivatePageHandler)
	})
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}
