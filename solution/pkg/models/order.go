package models

// OrderItem represents a single item within an order, containing the product reference and quantity.
type OrderItem struct {
	ProductID string `json:"productId"` // ID of the product being ordered
	Quantity  int    `json:"quantity"`  // Number of units of this product in the order
}

// OrderRequest represents the incoming request payload for creating a new order.
// It contains the items to be ordered and an optional coupon code for discounts.
type OrderRequest struct {
	CouponCode string      `json:"couponCode,omitempty"` // Optional coupon code for discounts
	Items      []OrderItem `json:"items"`                // List of items being ordered (required)
}

// Order represents a complete processed order with all calculated values.
// It includes the original request items plus calculated totals, applied discounts, and full product details.
type Order struct {
	ID        string      `json:"id"`        // Unique identifier for the order
	Total     float64     `json:"total"`     // Final total after discounts
	Discounts float64     `json:"discounts"` // Total discount amount applied
	Items     []OrderItem `json:"items"`     // Original items from the request
	Products  []Product   `json:"products"`  // Full product details for all ordered items
}
