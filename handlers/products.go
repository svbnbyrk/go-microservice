package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/sbvnbyrk/go-microservice/data"
)

//A list of products return in the response
//swagger:response productResponse
type productResponse struct{
	//All product at the system
	//in:body
	Body []data.Product
}

//swagger:response 	noContent
type productNoContent struct{

}
//swagger:parameters deleteProduct
type productIDParameterWrapper struct{
	//the id of the product to delete from database
	//in: path
	//required: true
	ID int `json:id`
}

//Products is a http handler
type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
			return
		}

		//validate the product
		err = prod.Validate()

		if err != nil {
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
