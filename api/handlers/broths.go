package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/carinamaleski/ramen-go/api/internal/models"
	"net/http"
	"os"
)

var broths = []models.Broth{
	{ID: 1, Name: "Salt", ImageInactive: "salt-inactive.svg", ImageActive: "salt-active.svg", Description: "Simple like the seawater, nothing more", Price: 10},
	{ID: 2, Name: "Shoyu", ImageInactive: "shoyu-inactive.svg", ImageActive: "shoyu-active.svg", Description: "The good old and traditional soy sauce", Price: 10},
	{ID: 3, Name: "Miso", ImageInactive: "miso-inactive.svg", ImageActive: "miso-active.svg", Description: "Paste made of fermented soybeans", Price: 12},
}

func GetBroths(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("x-api-key") != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "x-api-key header missing"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Sending broths data:", broths)
	json.NewEncoder(w).Encode(broths)
}
