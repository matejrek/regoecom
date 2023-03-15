package controllers

import (
	"database/sql"
	_ "encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-jwt/jwt/v4"
	_ "github.com/google/uuid"
	/*"golang.org/x/crypto/bcrypt"*/
	"net/http"
)

type Category struct {
	ID   int
	Name string
}

// GET all categories
func Categories(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT category FROM categories")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		// Prepare slice to hold categories
		var categories []string

		// Iterate over rows and add categories to slice
		var category string
		for rows.Next() {
			if err := rows.Scan(&category); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			categories = append(categories, category)
		}

		// Return categories as JSON
		c.JSON(http.StatusOK, categories)
	}
}
