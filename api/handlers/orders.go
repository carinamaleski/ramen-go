package handlers

import (
	"encoding/json"
	"github.com/carinamaleski/ramen-go/api/internal/models"
	"net/http"
	"os"
)

var orders []models.OrderResponse

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("x-api-key") != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "x-api-key header missing"})
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var orderRequest models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		http.Error(w, "Erro ao decodificar pedido", http.StatusBadRequest)
		return
	}

	if orderRequest.BrothID == 0 || orderRequest.ProteinID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "both brothId and proteinId are required"})
		return
	}

	orderID, err := generateOrderID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "could not place order"})
		return
	}

	newOrder := models.OrderResponse{
		ID:          orderID,
		Description: "Salt and Chasu Ramen",
		Image:       "public/images/lamen-miso.png",
	}

	orders = append(orders, newOrder)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}

func generateOrderID() (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.tech.redventures.com.br/orders/generate-id", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("x-api-key", "ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var response struct {
		OrderID string `json:"orderID"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.OrderID, nil
}
