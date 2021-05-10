package handlers

import (
	"Ex5Validation/data"
	"net/http"
)

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	productsList := data.GetProducts()
	err := productsList.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to encode json", http.StatusInternalServerError)
	}
}
