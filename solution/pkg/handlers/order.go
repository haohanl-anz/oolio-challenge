package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/haohanl/oolio-challenge/solution/pkg/models"
	"github.com/haohanl/oolio-challenge/solution/pkg/services"
)

// OrderHandler handles HTTP requests related to order operations.
type OrderHandler struct {
	orderService *services.OrderService
}

// NewOrderHandler creates a new OrderHandler with the provided OrderService.
func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// PlaceOrder handles POST /api/order requests.
// Processes order creation requests with items, optional coupon codes, and API key authentication.
// Returns the created order with calculated totals and applied discounts as JSON.
// Requires "api_key: apitest" header for authentication.
func (h *OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	// Check API key authentication
	apiKey := r.Header.Get("api_key")
	if apiKey != "apitest" {
		h.sendErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse request body
	var orderReq models.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		h.sendErrorResponse(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Create order
	order, err := h.orderService.CreateOrder(&orderReq)
	if err != nil {
		h.sendErrorResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *OrderHandler) sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResp := models.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	}

	json.NewEncoder(w).Encode(errorResp)
}
