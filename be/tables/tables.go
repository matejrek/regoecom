package tables

import (
	"database/sql"
	_ "encoding/json"
	_ "fmt"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-jwt/jwt/v4"
	_ "github.com/google/uuid"
	_ "golang.org/x/crypto/bcrypt"
	_ "log"
	_ "math/rand"
	"net/http"
	_ "sort"
	_ "strconv"
	_ "strings"
	_ "time"
)

// Generate Store tables
func StoreTables(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// SQL statement to create products table
		productsTable := `
	CREATE TABLE IF NOT EXISTS products (
	  id INT AUTO_INCREMENT PRIMARY KEY,
	  title VARCHAR(255) NOT NULL,
	  price DECIMAL(10, 2) NOT NULL,
	  currency VARCHAR(10) NOT NULL,
	  description TEXT NOT NULL
	);`

		// SQL statement to create categories table
		categoriesTable := `
	CREATE TABLE IF NOT EXISTS categories (
	  id INT AUTO_INCREMENT PRIMARY KEY,
	  category VARCHAR(255) NOT NULL
	);`

		// SQL statement to create product_categories junction table
		productCategoriesTable := `
	CREATE TABLE IF NOT EXISTS product_categories (
	  product_id INT NOT NULL,
	  category_id INT NOT NULL,
	  FOREIGN KEY (product_id) REFERENCES products(id),
	  FOREIGN KEY (category_id) REFERENCES categories(id)
	);`

		// Execute all SQL statements
		if _, err := db.Exec(productsTable); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if _, err := db.Exec(categoriesTable); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if _, err := db.Exec(productCategoriesTable); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return success message
		c.JSON(http.StatusOK, gin.H{"message": "Tables created successfully"})
	}
}

// Generate AUTH tables
func AuthTables(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS roles (id INT AUTO_INCREMENT PRIMARY KEY,	name VARCHAR(255) NOT NULL)`)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error 1": "database error"})
			return
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
				id INT AUTO_INCREMENT PRIMARY KEY,
				first_name VARCHAR(255) NOT NULL,
				last_name VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				password VARBINARY(255) NOT NULL,
				role_id INT NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
				FOREIGN KEY (role_id) REFERENCES roles(id)
		)`)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error 2": "database error"})
			return
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS refresh_tokens (
        id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT NOT NULL,
        token VARBINARY(255) NOT NULL,
        expires_at TIMESTAMP NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT fk_user_id
            FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
    )`)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error 3": "database error"})
			return
		}

		// Insert roles into roles table
		_, err = db.Exec(`INSERT INTO roles (name) VALUES ('admin'), ('user') ON DUPLICATE KEY UPDATE name = VALUES(name)`)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Tables generated successfully"})
	}
}
