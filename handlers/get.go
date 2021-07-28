package handlers

import (
	"net/http"

	"github.com/sbvnbyrk/go-microservice/data"
)

//swagger:route GET /products products listProduct
//Returns a list of product
//responses:
//	200:productResponse

//GetProduct return the product form the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	//fetch the products from datastore
	lp := data.GetProducts()

	//seriliaze list of JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable json writer", http.StatusInternalServerError)
	}
}
