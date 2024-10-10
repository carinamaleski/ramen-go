package main

import (
	"fmt"
	"net/http"
	"ramen-go/handlers"
)

func main() {
	http.HandleFunc("/broths", handlers.GetBroths)
	http.HandleFunc("/proteins", handlers.GetProteins)
	http.HandleFunc("/orders", handlers.CreateOrder)

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
