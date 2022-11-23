package handlers

import (
	"building-microservices-go/product-api/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Get all records")

	lp := data.GetProducts()
	err := data.ToJSON(lp, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
// responses:
//	200: productsResponse
//  404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Get single record")

	id := getProductID(r)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:
	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(err.Error(), w)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(err.Error(), w)
		return
	}

	err = data.ToJSON(prod, w)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
