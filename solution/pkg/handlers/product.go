package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/haohanl/oolio-challenge/solution/pkg/models"
	"github.com/haohanl/oolio-challenge/solution/pkg/services"
)

// ProductHandler handles HTTP requests related to product operations.
type ProductHandler struct {
	productService *services.ProductService
}

// NewProductHandler creates a new ProductHandler with the provided ProductService.
func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// ListProducts handles GET /api/product requests.
// Returns a JSON array of all available products with their details including
// ID, name, price, category, and image URLs for different screen sizes.
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products := h.productService.GetAllProducts()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// GetProduct handles GET /api/product/{id} requests.
// Extracts the product ID from the URL path and returns the product details as JSON.
// Returns 404 if the product is not found, 400 for invalid product ID format.
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	product, err := h.productService.GetProductByID(productID)
	if err != nil {
		h.sendErrorResponse(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandler) sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResp := models.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}

	json.NewEncoder(w).Encode(errorResp)
}
