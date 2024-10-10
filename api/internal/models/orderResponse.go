package models

type OrderResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
