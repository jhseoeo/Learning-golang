package main

import (
	"log"
	"net/http"
	"time"
)

func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
	})
}

var securityMsg = []byte("You didn't give the secret password\n")

func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(securityMsg)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func main() {
	terribleSecurity := TerribleSecurityProvider("GOPHER")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("greetings!\n"))
	})

	wrappedMux := terribleSecurity(RequestTimer(mux))

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  100 * time.Second,
		Handler:      wrappedMux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
