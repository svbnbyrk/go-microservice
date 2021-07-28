package handlers

import (
	"net/http"
	"github.com/sbvnbyrk/go-microservice/data"
)


func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Request")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}