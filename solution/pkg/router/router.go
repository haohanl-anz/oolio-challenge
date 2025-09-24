package router

import (
	"net/http"
	"strconv"

	"github.com/haohanl/oolio-challenge/solution/pkg/handlers"
	"github.com/haohanl/oolio-challenge/solution/pkg/middleware"
)

// Router handles HTTP routing
// It provides endpoints for product management and order processing with middleware support.
type Router struct {
	mux            *http.ServeMux
	productHandler *handlers.ProductHandler
	orderHandler   *handlers.OrderHandler
}

// New creates a new Router instance with the provided handlers.
// It initializes the HTTP ServeMux and sets up all route patterns.
func New(productHandler *handlers.ProductHandler, orderHandler *handlers.OrderHandler) *Router {
	r := &Router{
		mux:            http.NewServeMux(),
		productHandler: productHandler,
		orderHandler:   orderHandler,
	}

	r.setupRoutes()
	return r
}

// ServeHTTP implements the http.Handler interface.
// It applies middleware (logging and CORS) before routing requests to appropriate handlers.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Apply middleware chain to the mux
	handler := middleware.Chain(
		middleware.Logger,
		middleware.CORS,
	)(r.mux)

	handler.ServeHTTP(w, req)
}

func (r *Router) setupRoutes() {
	// API routes
	r.mux.HandleFunc("GET /api/product", r.handleProductList)
	r.mux.HandleFunc("GET /api/product/{id}", r.handleProductByID)
	r.mux.HandleFunc("POST /api/order", r.handleOrder)

	// Health check and service routes
	r.mux.HandleFunc("GET /api/health", r.healthCheck)
	r.mux.HandleFunc("GET /api", r.handleRoot)

	// Catch-all handler for unspecified paths - must be registered last
	r.mux.HandleFunc("/", r.handleNotFound)
}

func (r *Router) handleProductList(w http.ResponseWriter, req *http.Request) {
	r.productHandler.ListProducts(w, req)
}

func (r *Router) handleProductByID(w http.ResponseWriter, req *http.Request) {
	productID := req.PathValue("id")

	// Validate that the product ID is numeric
	if _, err := strconv.ParseInt(productID, 10, 64); err != nil {
		http.Error(w, "Invalid product ID format", http.StatusBadRequest)
		return
	}

	r.productHandler.GetProduct(w, req)
}

func (r *Router) handleOrder(w http.ResponseWriter, req *http.Request) {
	r.orderHandler.PlaceOrder(w, req)
}

func (r *Router) handleRoot(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Order Food Online API","version":"1.0.0","healthCheck":"/health"}`))
}

func (r *Router) healthCheck(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","message":"Server is healthy"}`))
}

func (r *Router) handleNotFound(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error":"Not Found","message":"The requested resource was not found"}`))
}
