package models

// ApiResponse represents a standard API response structure for general API communications.
// It follows common API response patterns with status code, type classification, and message.
type ApiResponse struct {
	Code    int    `json:"code"`    // HTTP status code
	Type    string `json:"type"`    // Response type classification
	Message string `json:"message"` // Human-readable response message
}

// ErrorResponse represents a structured error response for API endpoints.
// It provides consistent error formatting across all API endpoints.
type ErrorResponse struct {
	Error   string `json:"error"`   // Error type or category
	Message string `json:"message"` // Detailed error description
}
