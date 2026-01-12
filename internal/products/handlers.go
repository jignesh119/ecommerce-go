package products

import (
	"log"
	"net/http"

	"github.com/jignesh119/ecommerce-go/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// call the services, get products, convert into json, send it back to client

	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	products := []string{"bottle", "shirt"}
	json.Write(w, http.StatusOK, products)
}
