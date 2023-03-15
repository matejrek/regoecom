package controllers

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-jwt/jwt/v4"
	_ "github.com/google/uuid"
	/*"golang.org/x/crypto/bcrypt"*/
	"log"
	"net/http"
	"sort"
	"strconv"
)

// Product represents a product with its details
type Product struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Category    string     `json:"category"`
	Price       float64    `json:"price"`
	Currency    string     `json:"currency"`
	Description string     `json:"description"`
	Categories  []Category `json:"categories"`
}

// Create a new product
func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Bind JSON request body to Product struct
		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert product into database
		stmt, err := db.Prepare("INSERT INTO products (title, category, price, currency, description) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer stmt.Close()

		result, err := stmt.Exec(product.Title, product.Category, product.Price, product.Currency, product.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Get product ID
		productID, err := result.LastInsertId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return created product as JSON
		product.ID = int(productID)
		c.JSON(http.StatusCreated, product)
	}
}

// Get product by ID
func ProductById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get product ID from URL parameter
		productID := c.Param("id")

		// SQL statement to select product by ID
		stmt, err := db.Prepare("SELECT * FROM products WHERE id = ?")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Query the database for the product
		var product Product
		err = stmt.QueryRow(productID).Scan(&product.ID, &product.Title, &product.Price, &product.Currency, &product.Description)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		// Return the product as JSON
		c.JSON(http.StatusOK, product)
	}
}

// Get all products with categories
func Products(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Open a database connection
		db, err := sql.Open("mysql", "root:@/ecom")
		if err != nil {
			log.Fatalf("Could not connect to database: %v", err)
		}
		defer db.Close()

		// Query the database to get all products with their categories
		rows, err := db.Query(`
		SELECT products.id, products.Title, products.description, products.price, products.currency, categories.id, categories.category
		FROM products
		JOIN product_categories ON products.id = product_categories.product_id
		JOIN categories ON product_categories.category_id = categories.id
		ORDER BY products.id ASC
	`)
		if err != nil {
			log.Fatalf("Could not query database: %v", err)
		}
		defer rows.Close()

		// Create a map to store the products and their categories
		productMap := make(map[int]Product)
		for rows.Next() {
			var productID, categoryID int
			var productName, productDescription string
			var productCurrency string
			var productPrice float64
			var categoryName string

			err = rows.Scan(&productID, &productName, &productDescription, &productPrice, &productCurrency, &categoryID, &categoryName)
			if err != nil {
				log.Fatalf("Could not scan rows: %v", err)
			}

			// Get the product from the map, or create a new one if it doesn't exist
			product, ok := productMap[productID]
			if !ok {
				product = Product{
					ID:          productID,
					Title:       productName,
					Description: productDescription,
					Price:       productPrice,
					Currency:    productCurrency,
					Categories:  []Category{},
				}
			}

			// Append the category to the product's categories slice
			product.Categories = append(product.Categories, Category{
				ID:   categoryID,
				Name: categoryName,
			})

			// Store the product back in the map
			productMap[productID] = product
		}

		// Convert the map to a slice
		var products []Product
		for _, product := range productMap {
			products = append(products, product)
		}

		// Sort the products slice by ID
		sort.Slice(products, func(i, j int) bool {
			return products[i].ID < products[j].ID
		})

		// Return the response as JSON
		c.JSON(http.StatusOK, products)
	}
}

// Get products (max x)
func ProductsMaxLimit(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := c.Param("limit")
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}

		// Open a database connection
		db, err := sql.Open("mysql", "root:@/ecom")
		if err != nil {
			log.Fatalf("Could not connect to database: %v", err)
		}
		defer db.Close()

		// Query the database to get the specified number of products with their categories
		rows, err := db.Query(fmt.Sprintf(`
		SELECT products.id, products.Title, products.description, products.price, products.currency, categories.id, categories.category
		FROM products
		JOIN product_categories ON products.id = product_categories.product_id
		JOIN categories ON product_categories.category_id = categories.id
		ORDER BY products.id ASC
		LIMIT %d
	`, limitInt))
		if err != nil {
			log.Fatalf("Could not query database: %v", err)
		}
		defer rows.Close()

		// Create a map to store the products and their categories
		productMap := make(map[int]Product)
		for rows.Next() {
			var productID, categoryID int
			var productName, productDescription string
			var productCurrency string
			var productPrice float64
			var categoryName string

			err = rows.Scan(&productID, &productName, &productDescription, &productPrice, &productCurrency, &categoryID, &categoryName)
			if err != nil {
				log.Fatalf("Could not scan rows: %v", err)
			}

			// Get the product from the map, or create a new one if it doesn't exist
			product, ok := productMap[productID]
			if !ok {
				product = Product{
					ID:          productID,
					Title:       productName,
					Description: productDescription,
					Price:       productPrice,
					Currency:    productCurrency,
					Categories:  []Category{},
				}
			}

			// Append the category to the product's categories slice
			product.Categories = append(product.Categories, Category{
				ID:   categoryID,
				Name: categoryName,
			})

			// Store the product back in the map
			productMap[productID] = product
		}

		// Convert the map to a slice
		var products []Product
		for _, product := range productMap {
			products = append(products, product)
		}

		// Return the response as JSON
		c.JSON(http.StatusOK, products)
	}
}
