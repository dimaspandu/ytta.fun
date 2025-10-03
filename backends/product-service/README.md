# Product Service

The Product Service handles the lifecycle of products listed on the platform.
It ensures products follow the temporary publishing model and expire after their defined duration.

This implementation uses a **Go backend** with a **static JSON file** as the data source, providing a simple prototype suitable for local development and containerized deployment.

---

### Team Scope
- Provide APIs for product creation, publishing, and expiration.
- Store product details, metadata, and availability state.
- Support queries for browsing and filtering products.

---

## Features

* **RESTful API endpoints**:

  * `GET /products` : List all products
  * `GET /products/{id}` : Get single product by ID
  * `GET /health` : Health check endpoint
* **CORS enabled** to allow frontend applications to access the API.
* **Regex-based routing** without external libraries (no Gorilla Mux or other dependencies).
* **Static JSON data** stored in `data/products.json`.
* **Docker-ready** with multi-stage build for small runtime image.

---

## Product Model

Each product has the following fields:

| Field       | Type              | Description                  |
| ----------- | ----------------- | ---------------------------- |
| `id`        | int               | Unique product identifier    |
| `name`      | string            | Product name                 |
| `seller`    | string            | Seller name                  |
| `price`     | float64           | Product price                |
| `expiresAt` | string (ISO 8601) | Product expiration timestamp |
| `images`    | []string          | List of image URLs           |
| `variants`  | []string          | List of available variants   |

Example JSON:

```json
{
  "id": 1,
  "name": "Product A",
  "seller": "Seller 1",
  "price": 29.99,
  "expiresAt": "2025-10-03T12:00:00Z",
  "images": ["https://picsum.photos/400/300?random=1","https://picsum.photos/400/300?random=2"],
  "variants": ["Red", "Blue", "Green"]
}
```

---

## Getting Started

### Prerequisites

* Go 1.21+ installed
* Docker (optional, for containerized deployment)

### Run Locally

1. Clone the repository
2. Make sure `data/products.json` exists with your product data
3. Run the server:

```bash
go run main.go
```

4. Access endpoints:

   * `http://localhost:8080/products`
   * `http://localhost:8080/products/1`
   * `http://localhost:8080/health`

### Using Docker

1. Build the Docker image:

```bash
docker build -t product-service .
```

2. Run the container:

```bash
docker run -p 8080:8080 product-service
```

3. API is accessible at `http://localhost:8080`

---

## Team Scope

* Provide APIs for product creation, publishing, and expiration.
* Store product details, metadata, and availability state.
* Support queries for browsing and filtering products.
* Enable containerized deployment for testing and development.

---

## Notes

* Currently uses a static JSON file for simplicity; future improvements may include database support.
* Product expiration logic is static in the JSON; dynamic expiration handling can be added in future versions.
* CORS allows any origin by default, can be restricted to frontend domain in production.
