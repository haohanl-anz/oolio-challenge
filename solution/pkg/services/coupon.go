package services

import (
	"bufio"
	"os"
	"strings"
)

// CouponService manages coupon validation and discount calculations.
// It loads valid coupon codes from a file and provides methods to validate and apply discounts.
type CouponService struct {
	validCoupons map[string]bool
}

// NewCouponService creates a new CouponService instance and loads valid coupons from the specified file.
// If the file doesn't exist, the service will be created with an empty coupon list.
// Returns an error if the file exists but cannot be read properly.
func NewCouponService(validCouponsFilePath string) (*CouponService, error) {
	service := &CouponService{
		validCoupons: make(map[string]bool),
	}

	// Load valid coupons from file
	if err := service.loadValidCoupons(validCouponsFilePath); err != nil {
		return nil, err
	}

	return service, nil
}

func (s *CouponService) loadValidCoupons(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		// If file doesn't exist, just continue with empty coupon list
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coupon := strings.TrimSpace(scanner.Text())
		if coupon != "" {
			s.validCoupons[coupon] = true
		}
	}

	return scanner.Err()
}

// IsValidCoupon checks if the provided coupon code is valid.
// Returns true if the coupon code exists in the loaded valid coupons list.
func (s *CouponService) IsValidCoupon(code string) bool {
	return s.validCoupons[code]
}

// GetDiscountPercent returns the discount percentage for a given coupon code.
// Returns 0.0 if the coupon is invalid or doesn't exist.
// Currently, all valid coupons provide a 10% discount.
func (s *CouponService) GetDiscountPercent(code string) float64 {
	// Simple discount logic - in a real system this would be more complex
	if !s.IsValidCoupon(code) {
		return 0.0
	}

	// For this example, all valid coupons give 10% discount
	return 0.10
}
