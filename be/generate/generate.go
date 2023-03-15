package generate

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/google/uuid"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Generate 1000 products
func Products(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Random currencies
		currency := "EUR"

		// Random descriptions
		desc := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

		// Loop over the number of products to generate
		for i := 0; i < 1000; i++ {
			// Random product titles
			rand.Seed(time.Now().UnixNano())
			const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
			const (
				letterIdxBits = 6                    // 6 bits to represent a letter index
				letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
				letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
			)
			var src = rand.NewSource(time.Now().UnixNano())
			b := make([]byte, 10)
			for i, cache, remain := 9, src.Int63(), letterIdxMax; i >= 0; {
				if remain == 0 {
					cache, remain = src.Int63(), letterIdxMax
				}
				if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
					b[i] = letterBytes[idx]
					i--
				}
				cache >>= letterIdxBits
				remain--
			}
			// Random prices
			prices := []float64{rand.Float64()*390 + 10}

			// SQL statement to insert the product
			stmt, err := db.Prepare("INSERT INTO products (title, price, currency, description) VALUES (?, ?, ?, ?)")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Execute the SQL statement
			res, err := stmt.Exec(string(b), prices[0], currency, desc)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Get the last inserted ID
			productID, err := res.LastInsertId()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Random number of categories to assign to the product
			numCategories := rand.Intn(4) + 1
			for j := 0; j < numCategories; j++ {
				// Random category ID
				categoryID := rand.Intn(20) + 1

				// SQL statement to insert the product-category relationship
				stmt, err := db.Prepare("INSERT INTO product_categories (product_id, category_id) VALUES (?, ?)")
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				// Execute the SQL statement
				_, err = stmt.Exec(productID, categoryID)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "1000 random products generated successfully"})
	}
}

// Generate categories
func Categories(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Open database connection
		db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ecom")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer db.Close()

		// Call the function to generate categories
		categories := []string{"Clothing", "Electronics", "Home & Kitchen", "Beauty & Personal Care", "Sports & Outdoors", "Books", "Toys & Games", "Health & Household", "Tools & Home Improvement", "Pet Supplies", "Office Products", "Automotive", "Music", "Movies & TV", "Video Games", "Software", "Cell Phones & Accessories", "Camera & Photo", "Musical Instruments", "Grocery & Gourmet Food"}

		// Loop through the categories
		for i := 0; i < len(categories); i++ {
			// SQL statement to insert the category
			stmt, err := db.Prepare("INSERT INTO categories (category) VALUES (?)")
			if err != nil {
				log.Fatalf("Failed to prepare SQL statement: %v", err)
			}

			// Execute the SQL statement
			_, err = stmt.Exec(categories[i])
			if err != nil {
				log.Fatalf("Failed to execute SQL statement: %v", err)
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "20 categories generated successfully"})
	}
}
