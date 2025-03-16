package model

import (
	"testing"
	"time"
)

func TestNewReservation(t *testing.T) {
	// 準備
	now := time.Now()
	startTime := now.Add(1 * time.Hour)
	endTime := now.Add(2 * time.Hour)

	// 実行
	reservation := NewReservation(startTime, endTime)

	// 検証
	if reservation.ID == "" {
		t.Error("Expected ID to be set, got empty string")
	}

	if !reservation.StartTime.Equal(startTime) {
		t.Errorf("Expected StartTime to be %v, got %v", startTime, reservation.StartTime)
	}

	if !reservation.EndTime.Equal(endTime) {
		t.Errorf("Expected EndTime to be %v, got %v", endTime, reservation.EndTime)
	}

	if reservation.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set, got zero time")
	}

	if reservation.UpdatedAt.IsZero() {
		t.Error("Expected UpdatedAt to be set, got zero time")
	}
}