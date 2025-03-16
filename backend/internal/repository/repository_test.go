package repository

import (
	"testing"
	"time"

	"github.com/naofumi-fujii/489-yoyaku/backend/internal/model"
)

func TestInMemoryReservationRepository_Create(t *testing.T) {
	// 準備
	repo := NewInMemoryReservationRepository()
	now := time.Now()
	reservation := model.NewReservation(now, now.Add(1*time.Hour))

	// 実行
	err := repo.Create(reservation)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 保存されたか確認
	stored, err := repo.FindByID(reservation.ID)
	if err != nil {
		t.Errorf("Expected no error when finding reservation, got %v", err)
	}
	if stored == nil {
		t.Error("Expected to find reservation, got nil")
	}
}

func TestInMemoryReservationRepository_FindAll(t *testing.T) {
	// 準備
	repo := NewInMemoryReservationRepository()
	now := time.Now()
	reservation1 := model.NewReservation(now, now.Add(1*time.Hour))
	reservation2 := model.NewReservation(now.Add(2*time.Hour), now.Add(3*time.Hour))

	// 実行：予約を追加
	repo.Create(reservation1)
	repo.Create(reservation2)

	// 全件取得
	reservations, err := repo.FindAll()

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(reservations) != 2 {
		t.Errorf("Expected 2 reservations, got %d", len(reservations))
	}
}

func TestInMemoryReservationRepository_FindByID(t *testing.T) {
	// 準備
	repo := NewInMemoryReservationRepository()
	now := time.Now()
	reservation := model.NewReservation(now, now.Add(1*time.Hour))
	repo.Create(reservation)

	// テストケース
	tests := []struct {
		name    string
		id      string
		wantNil bool
	}{
		{
			name:    "存在するID",
			id:      reservation.ID,
			wantNil: false,
		},
		{
			name:    "存在しないID",
			id:      "non-existent-id",
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 実行
			got, err := repo.FindByID(tt.id)

			// 検証
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if (got == nil) != tt.wantNil {
				t.Errorf("FindByID() got = %v, want nil = %v", got, tt.wantNil)
			}
			if !tt.wantNil && got.ID != tt.id {
				t.Errorf("Expected ID %s, got %s", tt.id, got.ID)
			}
		})
	}
}

func TestInMemoryReservationRepository_Delete(t *testing.T) {
	// 準備
	repo := NewInMemoryReservationRepository()
	now := time.Now()
	reservation := model.NewReservation(now, now.Add(1*time.Hour))
	repo.Create(reservation)

	// 実行：存在するIDを削除
	err := repo.Delete(reservation.ID)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 削除されたか確認
	found, _ := repo.FindByID(reservation.ID)
	if found != nil {
		t.Error("Expected reservation to be deleted, but it still exists")
	}

	// 実行：存在しないIDを削除（エラーにならないことを確認）
	err = repo.Delete("non-existent-id")

	// 検証
	if err != nil {
		t.Errorf("Expected no error when deleting non-existent ID, got %v", err)
	}
}