package handlers

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/sbvnbyrk/go-microservice/data"
)

//swagger:route DELETE /products/{id} products deleteProduct
//Returns noContent
//responses:
//	201:noContent	
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!= nil{
		http.Error(rw,"Unable to convert id",http.StatusBadRequest)
		return
	}
	data.DeleteProduct(id)
}