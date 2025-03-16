package handler

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/service"
)

type ReservationHandler struct {
	service  *service.ReservationService
	validate *validator.Validate
}

func NewReservationHandler(service *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		service:  service,
		validate: validator.New(),
	}
}

type createReservationRequest struct {
	StartTime string `json:"startTime" validate:"required"`
	EndTime   string `json:"endTime" validate:"required"`
}

func (h *ReservationHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/api/reservations", h.CreateReservation)
	e.GET("/api/reservations", h.GetAllReservations)
	e.DELETE("/api/reservations/:id", h.DeleteReservation)
}

func (h *ReservationHandler) CreateReservation(c echo.Context) error {
	req := new(createReservationRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid start time format"})
	}

	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid end time format"})
	}

	if endTime.Before(startTime) || endTime.Equal(startTime) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "End time must be after start time"})
	}

	params := service.CreateReservationParams{
		StartTime: startTime,
		EndTime:   endTime,
	}

	reservation, err := h.service.CreateReservation(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create reservation"})
	}

	return c.JSON(http.StatusCreated, reservation)
}

func (h *ReservationHandler) GetAllReservations(c echo.Context) error {
	reservations, err := h.service.GetAllReservations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get reservations"})
	}

	return c.JSON(http.StatusOK, reservations)
}

func (h *ReservationHandler) DeleteReservation(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	if err := h.service.DeleteReservation(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete reservation"})
	}

	return c.NoContent(http.StatusNoContent)
}
