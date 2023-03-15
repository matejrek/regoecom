package main

import (
	"be/auth"
	"be/controllers"
	"be/generate"
	"be/tables"
	"database/sql"
	_ "encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-jwt/jwt/v4"
	_ "github.com/google/uuid"
	/*"golang.org/x/crypto/bcrypt"*/
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database credentials from the environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Connect to the database using the credentials from the ENV file
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, dbname))

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// Create Gin router
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization") // add authorization header to allowed headers

	router.Use(cors.New(config))

	// Route to create a new product
	router.POST("/products", controllers.CreateProduct(db))

	// Route to get product by ID
	router.GET("/products/:id", controllers.ProductById(db))

	// Route to get all products with categories
	router.GET("/products", controllers.Products(db))

	// Route to get products with max limit of how many
	router.GET("/products/max/:limit", controllers.ProductsMaxLimit(db))

	// Route to get all categories
	router.GET("/categories", controllers.Categories(db))

	/***********************************************
	  ************************************************
	  STORE TABLE & DUMMY DATA GENERATION
	  ************************************************
	  ***********************************************/

	// Route to create store tables
	router.GET("/create-tables", tables.StoreTables(db))

	// Route to generate random categories
	router.GET("/generate/categories", generate.Categories(db))

	// Route to generate 1000 random products
	router.GET("/generate/products", generate.Products(db))

	/***********************************************
	  ************************************************
	  AUTH ROUTES & AUTH TABLES
	  ************************************************
	  ***********************************************/

	// Registration route
	router.POST("/register", auth.Register(db))
	// Login route
	router.POST("/login", auth.Login(db))
	// User route
	router.GET("/user", auth.GetUser(db))
	// Generate auth tables
	router.GET("/generate-auth-tables", tables.AuthTables(db))

	// Start server
	host := "localhost"
	port := "9010"
	if err := router.Run(host + ":" + port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
