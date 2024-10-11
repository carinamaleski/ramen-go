package main

import (
	"fmt"
	"github.com/carinamaleski/ramen-go/api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/broths", enableCORS(handlers.GetBroths))
	http.HandleFunc("/proteins", enableCORS(handlers.GetProteins))
	http.HandleFunc("/orders", enableCORS(handlers.CreateOrder))

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-api-key")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}
