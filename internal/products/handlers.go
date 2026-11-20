package products

import (
	"net/http"

	"github.com/sikozonpc/ecom/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Listproducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.Listproducts(r.Context())
	if err != nil {
		json.Write(w, http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{Error: "Failed to list products"})
		return
	}

	json.Write(w, http.StatusOK, struct {
		Products []string `json:"products"`
	}{Products: products})
}
