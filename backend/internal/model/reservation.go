package model

import (
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID        string    `json:"id"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewReservation(startTime, endTime time.Time) *Reservation {
	now := time.Now()
	return &Reservation{
		ID:        uuid.New().String(),
		StartTime: startTime,
		EndTime:   endTime,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
