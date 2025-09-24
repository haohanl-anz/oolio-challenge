package services

import (
	"errors"
	"fmt"

	"github.com/haohanl/oolio-challenge/solution/pkg/models"
)

// OrderService handles order processing logic including validation, pricing, and discount application.
// It coordinates between ProductService and CouponService to create complete orders.
type OrderService struct {
	productService *ProductService
	couponService  *CouponService
}

// NewOrderService creates a new OrderService with the required dependencies.
func NewOrderService(productService *ProductService, couponService *CouponService) *OrderService {
	return &OrderService{
		productService: productService,
		couponService:  couponService,
	}
}

// CreateOrder processes an order request and returns a complete order with calculated totals.
// Validates the request, retrieves products, calculates subtotal, applies coupon discounts if valid,
// and generates a unique order ID. Returns an error if validation fails or products are not found.
func (s *OrderService) CreateOrder(req *models.OrderRequest) (*models.Order, error) {
	// Validate request
	if len(req.Items) == 0 {
		return nil, errors.New("order must contain at least one item")
	}

	// Extract product IDs and validate quantities
	var productIDs []string
	for _, item := range req.Items {
		if item.ProductID == "" {
			return nil, errors.New("product ID is required for all items")
		}
		if item.Quantity <= 0 {
			return nil, errors.New("quantity must be greater than 0")
		}
		productIDs = append(productIDs, item.ProductID)
	}

	// Get products
	products, err := s.productService.GetProductsByIDs(productIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products: %w", err)
	}

	// Calculate subtotal
	productMap := make(map[string]models.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	var subtotal float64
	for _, item := range req.Items {
		product, exists := productMap[item.ProductID]
		if !exists {
			return nil, fmt.Errorf("product with ID %s not found", item.ProductID)
		}
		subtotal += product.Price * float64(item.Quantity)
	}

	// Apply coupon discount if provided
	var discountAmount float64
	if req.CouponCode != "" {
		if s.couponService.IsValidCoupon(req.CouponCode) {
			discountPercent := s.couponService.GetDiscountPercent(req.CouponCode)
			discountAmount = subtotal * discountPercent
		}
	}

	total := subtotal - discountAmount

	// Generate order ID (in a real system, this would be more sophisticated)
	orderID := s.generateOrderID()

	order := &models.Order{
		ID:        orderID,
		Total:     total,
		Discounts: discountAmount,
		Items:     req.Items,
		Products:  products,
	}

	return order, nil
}

func (s *OrderService) generateOrderID() string {
	// Simple UUID-like generation for demo purposes
	// In a real system, you'd use a proper UUID library or database sequence
	return fmt.Sprintf("%04d-%04d-%04d-%04d",
		1000+len(s.productService.GetAllProducts()),
		2000,
		3000,
		4000)
}
