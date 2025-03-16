package main

import (
	"net/http"

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

	// Initialize repository
	reservationRepo := repository.NewInMemoryReservationRepository()

	// Initialize service
	reservationService := service.NewReservationService(reservationRepo)

	// Initialize handler
	reservationHandler := handler.NewReservationHandler(reservationService)

	// Register routes
	reservationHandler.RegisterRoutes(e)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
