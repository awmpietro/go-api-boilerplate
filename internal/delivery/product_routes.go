package delivery

import (
	"net/http"
)

func RegisterProductRoutes(productHandler ProductHandler, mux *http.ServeMux) {
	mux.HandleFunc("/products", productHandler.GetProducts)
	mux.HandleFunc("/products/create", productHandler.CreateProduct)
	mux.HandleFunc("/products/update", productHandler.UpdateProduct)
	mux.HandleFunc("/products/delete", productHandler.DeleteProduct)
}
