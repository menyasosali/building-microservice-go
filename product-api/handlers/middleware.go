package handlers

import (
	"building-microservices-go/product-api/data"
	"compress/gzip"
	"context"
	"fmt"
	"net/http"
)

func GZipResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// write the file
		next.ServeHTTP(w, r)
	})
}

type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product
		err = data.Validate(prod)
		if err != nil {
			p.l.Println("[ERROR] Validating product", err)
			http.Error(w, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
