package handlers

import (
	"building-microservices-go/product-api/data"
	"net/http"
)

// Create handles POST requests to add new products
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}
