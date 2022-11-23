package data

import (
	"fmt"
	"time"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for this user
	//
	// required: true
	// min: 1
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"required,gt=0"`
	SKU         string  `json:"SKU" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}

var ErrProductNotFound = fmt.Errorf("Product not found")

// AddProduct adds a new product to the database
func AddProduct(p *Product) {
	// get the next id in sequence
	maxID := productList[len(productList)-1]
	p.ID = maxID.ID + 1
	productList = append(productList, p)
}

// GetProducts returns all products from the database
func GetProducts() Products {
	return productList
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	pos := findIndexByProductId(id)
	if pos == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:pos], productList[:pos+1]...)
	return nil
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func UpdateProduct(product *Product) error {
	pos := findIndexByProductId(product.ID)
	if pos == -1 {
		return ErrProductNotFound
	}

	// update the product in the DB
	productList[pos] = product

	return nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductId(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*Product, error) {
	pos := findIndexByProductId(id)
	if pos == -1 {
		return nil, ErrProductNotFound
	}
	return productList[pos], nil
}
