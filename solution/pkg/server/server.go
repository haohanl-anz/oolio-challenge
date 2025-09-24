package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/haohanl/oolio-challenge/solution/pkg/handlers"
	"github.com/haohanl/oolio-challenge/solution/pkg/router"
	"github.com/haohanl/oolio-challenge/solution/pkg/services"
)

// Server represents an HTTP server with graceful shutdown capabilities.
type Server struct {
	httpServer *http.Server
	address    string
}

// New creates a new Server instance with the specified address.
// The address should be in the format ":port" or "host:port".
func New(address string) *Server {
	return &Server{
		address: address,
	}
}

// Start initializes all services and handlers, then starts the HTTP server.
// It blocks until the server is gracefully shut down via interrupt signals.
// Returns an error if the server fails to initialize or start.
func (s *Server) Start() error {
	log.Println("Initializing services...")

	// Initialize services
	productService := services.NewProductService()

	couponService, err := services.NewCouponService("coupons/valid_coupons.txt")
	if err != nil {
		log.Printf("Warning: Could not load coupon service: %v", err)
		couponService, _ = services.NewCouponService("") // Create empty service
	}

	orderService := services.NewOrderService(productService, couponService)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Initialize router
	appRouter := router.New(productHandler, orderHandler)

	// Create HTTP server
	s.httpServer = &http.Server{
		Addr:         s.address,
		Handler:      appRouter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Starting server on %s", s.address)

	// Start server in goroutine
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	s.waitForShutdown()

	return nil
}

// waitForShutdown blocks until an interrupt signal (SIGINT or SIGTERM) is received,
// then initiates a graceful shutdown of the HTTP server with a 30-second timeout.
func (s *Server) waitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server is shutting down...")

	// Create a context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
