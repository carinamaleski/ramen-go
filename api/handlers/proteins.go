package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"ramen-go/internal/models"
)

var proteins = []models.Protein{
	{ID: 1, ImageInactive: "inactive1.png", ImageActive: "active1.png", Name: "Chasu", Description: "A sliced flavourful pork meat with a selection of season vegetables.", Price: 10},
	{ID: 2, ImageInactive: "inactive2.png", ImageActive: "active2.png", Name: "Yasai Vegetarian", Description: "A delicious vegetarian lamen with a selection of season vegetables.", Price: 10},
	{ID: 3, ImageInactive: "inactive3.png", ImageActive: "active3.png", Name: "Karaague", Description: "Three units of fried chicken, moyashi, ajitama egg and other vegetables.", Price: 12}}

func GetProteins(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("x-api-key") != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "x-api-key header missing"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proteins)
}
