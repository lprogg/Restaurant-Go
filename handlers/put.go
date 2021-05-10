package handlers

import (
	"Ex5Validation/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *Products) PutProducts(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Products", id)

	product := r.Context().Value(&KeyProduct{}).(*data.Product)

	err = product.PutProduct(id)
	if err == data.ErrorProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
