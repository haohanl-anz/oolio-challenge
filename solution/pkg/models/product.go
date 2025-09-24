package models

// Product represents a product available for order in the system.
// It contains all the necessary information for display and ordering including
// pricing, categorization, and responsive image URLs.
type Product struct {
	ID       string        `json:"id"`       // Unique identifier for the product
	Name     string        `json:"name"`     // Display name of the product
	Price    float64       `json:"price"`    // Price of the product in the system currency
	Category string        `json:"category"` // Product category for organization
	Image    ProductImages `json:"image"`    // Responsive image URLs for different screen sizes
}

// ProductImages contains image URLs optimized for different screen sizes and devices.
// This supports responsive design by providing appropriate image resolutions.
type ProductImages struct {
	Thumbnail string `json:"thumbnail"` // Small image for thumbnails and previews
	Mobile    string `json:"mobile"`    // Image optimized for mobile devices
	Tablet    string `json:"tablet"`    // Image optimized for tablet devices
	Desktop   string `json:"desktop"`   // High-resolution image for desktop displays
}
