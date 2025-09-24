# Order Food Online - Go Microservice

A RESTful API microservice built with Go's standard library.
Please reach out for further questions regarding this implementation.

## Architecture

```
solution/
├── cmd/
│   ├── server/         # Main application entry point
│   └── coupons/        # Coupon validation utility
├── pkg/
│   ├── handlers/       # HTTP request handlers
│   ├── middleware/     # HTTP middleware (CORS, logging)
│   ├── models/         # Data structures and DTOs
│   ├── router/         # HTTP routing with Go 1.22+ patterns
│   ├── server/         # Server initialization and lifecycle
│   └── services/       # Business logic layer
├── coupons/           # Coupon data files
└── go.mod             # Go module definition
```

## Prerequisites

- Go 1.22 or higher
- Access to terminal/command line

## Quick Start

### 1. Initialize the Project

Clone or navigate to the solution directory:

```bash
cd /path/to/oolio-challenge/solution
```

### 2. Generate Valid Coupons (Optional)

If the coupon base files are present in `.coupons/`, the following script can be executed to find valid coupons:

```bash
go run cmd/coupons/main.go
```

This will output a `valid_coupons.txt` file which contains validated coupons which show up in at least 2 of the 3 provided coupon base files. 

### 3. Start the Server

```bash
go run cmd/server/main.go
```

The server will start on port 8080 by default. You should see output like:

```
Initializing services...
Starting server on :8080
```

### 5. Test the API

Check if the server is running:

```bash
curl http://localhost:8080/api/health
```

Expected response:
```json
{
  "status": "ok",
  "message": "Server is healthy"
}
```

## API Endpoints

### Base URL
```
http://localhost:8080
```

### Available Endpoints

| Method | Endpoint            | Description       | Auth Required |
| ------ | ------------------- | ----------------- | ------------- |
| GET    | `/api`              | API information   | No            |
| GET    | `/api/health`       | Health check      | No            |
| GET    | `/api/product`      | List all products | No            |
| GET    | `/api/product/{id}` | Get product by ID | No            |
| POST   | `/api/order`        | Place an order    | Yes (API Key) |

### Authentication

The order endpoint requires API key authentication:

```bash
curl -H "api_key: apitest" -X POST http://localhost:8080/api/order
```

## API Usage Examples

### 1. List All Products

```bash
curl http://localhost:8080/api/product
```

Response:
```json
[
  {
    "id": "1",
    "name": "Vanilla Bean Crème Brûlée",
    "price": 7.0,
    "category": "Crème Brûlée",
    "image": {
      "thumbnail": "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-thumbnail.jpg",
      "mobile": "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-mobile.jpg",
      "tablet": "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-tablet.jpg",
      "desktop": "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-desktop.jpg"
    }
  }
]
```

### 2. Get Product by ID

```bash
curl http://localhost:8080/api/product/1
```

### 3. Place an Order

```bash
curl -X POST http://localhost:8080/api/order \
  -H "Content-Type: application/json" \
  -H "api_key: apitest" \
  -d '{
    "items": [
      {
        "productId": "1",
        "quantity": 2
      },
      {
        "productId": "2", 
        "quantity": 1
      }
    ],
    "couponCode": "HAPPYHRS"
  }'
```

Response:
```json
{
  "id": "1009-2000-3000-4000",
  "total": 19.8,
  "discounts": 2.2,
  "items": [
    {
      "productId": "1",
      "quantity": 2
    },
    {
      "productId": "2",
      "quantity": 1
    }
  ],
  "products": [...]
}
```

## Configuration

### Server Port

To run on a different port, modify `cmd/server/main.go`:

```go
srv := server.New(":3000") // Change from :8080
```

### Coupon File Path

The coupon service looks for `coupons/valid_coupons.txt` by default. To change this, modify the path in `pkg/server/server.go`:

```go
couponService, err := services.NewCouponService("path/to/your/coupons.txt")
```

## Error Responses

The API returns structured JSON error responses:

### 400 Bad Request
```json
{
  "error": "Bad Request",
  "message": "Invalid product ID format"
}
```

### 401 Unauthorized  
```json
{
  "error": "Unauthorized",
  "message": "Unauthorized"
}
```

### 404 Not Found
```json
{
  "error": "Not Found", 
  "message": "The requested resource was not found"
}
```

### 405 Method Not Allowed
```json
{
  "error": "Method Not Allowed",
  "message": "Method not allowed"
}
```

## Middleware

The application includes built-in middleware:

- **Logger**: Logs all HTTP requests with method, URI, client address, and duration
- **CORS**: Enables cross-origin requests with appropriate headers

### Adding Custom Middleware

To add custom middleware, modify the chain in `pkg/router/router.go`:

```go
handler := middleware.Chain(
    middleware.Logger,
    middleware.CORS,
    yourCustomMiddleware, // Add here
)(r.mux)
```
