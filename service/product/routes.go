package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/petersizovdev/go-ecommerce/types"
	"github.com/petersizovdev/go-ecommerce/utils"
)

type Handler struct{
	store types.ProductStore

}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) handleCreatedProduct(w http.ResponseWriter, r *http.Request){
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.Handle("/products", http.HandlerFunc(h.handleCreatedProduct)).Methods(http.MethodDelete)
}