package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/handler"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/repository"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/service"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Database connection
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "reservations")

	// Log DB connection info for debugging
	fmt.Printf("Connecting to MySQL at %s:%s as %s...\n", dbHost, dbPort, dbUser)
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci", 
		dbUser, dbPassword, dbHost, dbPort, dbName)
	
	// Initialize DB connection with retry logic
	var db *sql.DB
	var err error
	
	// More aggressive retry logic
	maxRetries := 60
	fmt.Printf("Attempting database connection (will retry %d times)...\n", maxRetries)
	
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("Successfully connected to the database!")
				break
			}
		}
		
		fmt.Printf("Failed to connect to database (attempt %d/%d): %v\n", i+1, maxRetries, err)
		time.Sleep(5 * time.Second)
	}
	
	if err != nil {
		e.Logger.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
	}
	
	// Set database connection parameters
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	
	// Initialize repository
	var reservationRepo repository.ReservationRepository
	
	mysqlRepo, err := repository.NewMySQLReservationRepository(db)
	if err != nil {
		e.Logger.Fatalf("Failed to initialize MySQL repository: %v", err)
	}
	reservationRepo = mysqlRepo

	// Initialize service
	reservationService := service.NewReservationService(reservationRepo)

	// Initialize handler
	reservationHandler := handler.NewReservationHandler(reservationService)

	// Register routes
	reservationHandler.RegisterRoutes(e)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok", 
			"db":     "connected",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// getEnv gets environment variable or returns default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
