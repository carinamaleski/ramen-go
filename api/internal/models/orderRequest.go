package models

type OrderRequest struct {
	BrothID   int `json:"brothId,string,omitempty"`
	ProteinID int `json:"proteinId,string,omitempty"`
}
