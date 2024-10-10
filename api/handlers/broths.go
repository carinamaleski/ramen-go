package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"ramen-go/internal/models"
)

var broths = []models.Broth{
	{ID: 1, ImageInactive: "inactive1.png", ImageActive: "active1.png", Name: "Salt", Description: "Simple like the seawater, nothing more", Price: 10},
	{ID: 2, ImageInactive: "inactive2.png", ImageActive: "active2.png", Name: "Shoyu", Description: "The good old and traditional soy sauce", Price: 10},
	{ID: 3, ImageInactive: "inactive3.png", ImageActive: "active3.png", Name: "Miso", Description: "Paste made of fermented soybeans", Price: 12},
}

func GetBroths(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("x-api-key") != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "x-api-key header missing"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(broths)
}
