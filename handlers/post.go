package handlers

import (
	"Ex5Validation/data"
	"net/http"
)

func (p *Products) PostProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	product := r.Context().Value(&KeyProduct{}).(*data.Product)
	product.InsertProduct()
}
