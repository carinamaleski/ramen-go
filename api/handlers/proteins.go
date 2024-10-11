package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/carinamaleski/ramen-go/api/internal/models"
	"net/http"
	"os"
)

var proteins = []models.Protein{
	{ID: 1, Name: "Chasu", ImageInactive: "pork-inactive.svg", ImageActive: "pork-active.svg", Description: "A sliced flavourful pork meat with a selection of season vegetables.", Price: 10},
	{ID: 2, Name: "Yasai Vegetarian", ImageInactive: "yasai-inactive.svg", ImageActive: "yasai-active.svg", Description: "A delicious vegetarian lamen with a selection of season vegetables.", Price: 10},
	{ID: 3, Name: "Karaague", ImageInactive: "chicken-inactive.svg", ImageActive: "chicken-active.svg", Description: "Three units of fried chicken, moyashi, ajitama egg and other vegetables.", Price: 12}}

func GetProteins(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("x-api-key") != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "x-api-key header missing"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Sending proteins data:", proteins)
	json.NewEncoder(w).Encode(proteins)
}
