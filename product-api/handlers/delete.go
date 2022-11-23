package handlers

import (
	"building-microservices-go/product-api/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
//	201: noContent

// DeleteProduct handles DELETE requests and removes items from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(err.Error(), w)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(err.Error(), w)
		return
	}
}
