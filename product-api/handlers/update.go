package handlers

import (
	"building-microservices-go/product-api/data"
	"net/http"
)

func (p *Products) Update(w http.ResponseWriter, r *http.Request) {

	id := getProductID(r)

	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(&prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
