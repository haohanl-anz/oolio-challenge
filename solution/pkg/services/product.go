package services

import (
	"errors"
	"fmt"

	"github.com/haohanl/oolio-challenge/solution/pkg/models"
)

// ProductService provides business logic for product operations.
// It manages product data and provides methods for retrieval and lookup.
type ProductService struct {
	products []models.Product
}

// NewProductService creates a new ProductService instance initialized with sample product data.
// In a production environment, this would typically load data from a database.
func NewProductService() *ProductService {
	// Initialize with some sample products
	products := []models.Product{
		{
			ID:       "1",
			Name:     "Vanilla Bean Crème Brûlée",
			Price:    7.0,
			Category: "Crème Brûlée",
			Image: models.ProductImages{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-desktop.jpg",
			},
		},
		{
			ID:       "2",
			Name:     "Vanilla Macaron",
			Price:    8.0,
			Category: "Macaron",
			Image: models.ProductImages{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-macaron-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-macaron-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-macaron-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-macaron-desktop.jpg",
			},
		},
		{
			ID:       "3",
			Name:     "Classic Tiramisu",
			Price:    5.5,
			Category: "Tiramisu",
			Image: models.ProductImages{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-tiramisu-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-tiramisu-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-tiramisu-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-tiramisu-desktop.jpg",
			},
		},
		{
			ID:       "4",
			Name:     "Pistachio Baklava",
			Price:    4.0,
			Category: "Baklava",
			Image: models.ProductImages{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-baklava-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-baklava-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-baklava-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-baklava-desktop.jpg",
			},
		},
		{
			ID:       "5",
			Name:     "Lemon Meringue Pie",
			Price:    5.0,
			Category: "Pie",
			Image: models.ProductImages{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-meringue-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-meringue-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-meringue-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-meringue-desktop.jpg",
			},
		},
		{
			ID:       "6",
			Name:     "Red Velvet Cake",
			Price:    4.5,
			Category: "Cake",
			Image: models.ProductImages{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-cake-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-cake-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-cake-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-cake-desktop.jpg",
			},
		},
		{
			ID:       "7",
			Name:     "Salted Caramel Brownie",
			Price:    5.5,
			Category: "Brownie",
			Image: models.ProductImages{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-brownie-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-brownie-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-brownie-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-brownie-desktop.jpg",
			},
		},
		{
			ID:       "8",
			Name:     "Vanilla Panna Cotta",
			Price:    6.5,
			Category: "Panna Cotta",
			Image: models.ProductImages{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-desktop.jpg",
			},
		},
		{
			ID:       "9",
			Name:     "Classic Waffle",
			Price:    13.3,
			Category: "Waffle",
			Image: models.ProductImages{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
			},
		},
	}

	return &ProductService{
		products: products,
	}
}

// GetAllProducts returns all available products.
// Returns a slice containing all products in the system.
func (s *ProductService) GetAllProducts() []models.Product {
	return s.products
}

// GetProductByID retrieves a product by its ID.
// Returns the product if found, or an error if the product does not exist.
func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	for _, product := range s.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

// GetProductsByIDs retrieves multiple products by their IDs.
// Returns a slice of found products and an error if any products are not found.
// The error will list all missing product IDs.
func (s *ProductService) GetProductsByIDs(ids []string) ([]models.Product, error) {
	var products []models.Product
	var notFound []string

	for _, id := range ids {
		product, err := s.GetProductByID(id)
		if err != nil {
			notFound = append(notFound, id)
			continue
		}
		products = append(products, *product)
	}

	if len(notFound) > 0 {
		return products, fmt.Errorf("products not found: %v", notFound)
	}

	return products, nil
}
