package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"ecommerce/db"
	"ecommerce/models"
	"ecommerce/repository"
)

func main() {
	// Set up a context with a timeout for startup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Starting e-commerce backend service...")

	// Initialize the database connection and schema
	pool, err := db.InitDB(ctx)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer pool.Close()

	// Create repository
	productRepo := repository.NewProductRepository(pool)

	// Create a new context for runtime queries
	queryCtx := context.Background()

	// Insert a test product
	testProduct := &models.Product{
		Name:        "Ultimate mechanical keyboard",
		Description: "RGB backlit mechanical gaming keyboard with linear red switches.",
		Price:       129.99,
		Stock:       25,
	}

	log.Printf("Inserting test product: %s...", testProduct.Name)
	err = productRepo.Create(queryCtx, testProduct)
	if err != nil {
		log.Fatalf("Failed to create product: %v", err)
	}
	log.Printf("Successfully inserted product. ID: %d, CreatedAt: %s", testProduct.ID, testProduct.CreatedAt.Format(time.RFC3339))

	// Fetch all products
	log.Println("Retrieving products from database...")
	products, err := productRepo.GetAll(queryCtx)
	if err != nil {
		log.Fatalf("Failed to retrieve products: %v", err)
	}

	log.Printf("Found %d product(s) in database:", len(products))
	for _, p := range products {
		fmt.Printf("- [%d] %s (Price: $%.2f, Stock: %d) - %s\n", p.ID, p.Name, p.Price, p.Stock, p.Description)
	}
}