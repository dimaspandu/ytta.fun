package main

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
  "regexp"
  "strconv"
)

/*
Simple RESTful Product Service in Go

Endpoints:
- GET /products           : List all products
- GET /products/{id}      : Get single product by ID
- GET /health             : Health check

Features:
- CORS enabled for all origins
- Regex-based routing (no external libraries)
- Products loaded from 'data/products.json'
*/

type Product struct {
  ID        int      `json:"id"`
  Name      string   `json:"name"`
  Seller    string   `json:"seller"`
  Price     float64  `json:"price"`
  ExpiresAt string   `json:"expiresAt"`
  Images    []string `json:"images"`
  Variants  []string `json:"variants"`
}

var products []Product

func loadProducts() {
  data, err := ioutil.ReadFile("data/products.json")
  if err != nil {
    log.Fatalf("Error reading JSON file: %v", err)
  }
  if err := json.Unmarshal(data, &products); err != nil {
    log.Fatalf("Error unmarshalling JSON: %v", err)
  }
}

// CORS middleware
func withCORS(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    if r.Method == http.MethodOptions {
      w.WriteHeader(http.StatusOK)
      return
    }
    next.ServeHTTP(w, r)
  })
}

// Router handles requests with regex
func router(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path
  method := r.Method

  // /health route
  if path == "/health" && method == "GET" {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
    return
  }

  // /products route
  if path == "/products" && method == "GET" {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
    return
  }

  // /products/{id} route using regex
  re := regexp.MustCompile(`^/products/(\d+)$`)
  matches := re.FindStringSubmatch(path)
  if len(matches) == 2 && method == "GET" {
    id, _ := strconv.Atoi(matches[1])
    for _, p := range products {
      if p.ID == id {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(p)
        return
      }
    }
    http.Error(w, "Product not found", http.StatusNotFound)
    return
  }

  // If no route matched
  http.NotFound(w, r)
}

func main() {
  loadProducts()
  http.Handle("/", withCORS(http.HandlerFunc(router)))
  log.Println("Server running at :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
