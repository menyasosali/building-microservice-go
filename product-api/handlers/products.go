// Package classification of Product API
//
// Documentation for Product API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
// 	- application/json
//
//	Produces:
// 	- application/json
// swagger:meta

package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct{}

func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
