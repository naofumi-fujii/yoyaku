package service

import (
	"testing"
	"time"

	"github.com/naofumi-fujii/489-yoyaku/backend/internal/model"
)

// モックリポジトリの実装
type mockReservationRepository struct {
	reservations map[string]*model.Reservation
	createFunc   func(reservation *model.Reservation) error
	findAllFunc  func() ([]*model.Reservation, error)
	findByIDFunc func(id string) (*model.Reservation, error)
	deleteFunc   func(id string) error
}

func newMockReservationRepository() *mockReservationRepository {
	return &mockReservationRepository{
		reservations: make(map[string]*model.Reservation),
		createFunc: func(reservation *model.Reservation) error {
			return nil
		},
		findAllFunc: func() ([]*model.Reservation, error) {
			return []*model.Reservation{}, nil
		},
		findByIDFunc: func(id string) (*model.Reservation, error) {
			return nil, nil
		},
		deleteFunc: func(id string) error {
			return nil
		},
	}
}

func (m *mockReservationRepository) Create(reservation *model.Reservation) error {
	return m.createFunc(reservation)
}

func (m *mockReservationRepository) FindAll() ([]*model.Reservation, error) {
	return m.findAllFunc()
}

func (m *mockReservationRepository) FindByID(id string) (*model.Reservation, error) {
	return m.findByIDFunc(id)
}

func (m *mockReservationRepository) Delete(id string) error {
	return m.deleteFunc(id)
}

func TestReservationService_CreateReservation(t *testing.T) {
	// 準備
	mockRepo := newMockReservationRepository()
	var savedReservation *model.Reservation

	mockRepo.createFunc = func(reservation *model.Reservation) error {
		savedReservation = reservation
		return nil
	}

	service := NewReservationService(mockRepo)
	now := time.Now()
	params := CreateReservationParams{
		StartTime: now,
		EndTime:   now.Add(1 * time.Hour),
	}

	// 実行
	createdReservation, err := service.CreateReservation(params)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if createdReservation == nil {
		t.Fatal("Expected reservation to be created, got nil")
	}
	if savedReservation == nil {
		t.Fatal("Expected reservation to be saved in repository, but it was not")
	}
	if createdReservation.ID != savedReservation.ID {
		t.Errorf("Expected created reservation ID %s to match saved reservation ID %s", createdReservation.ID, savedReservation.ID)
	}
	if !createdReservation.StartTime.Equal(params.StartTime) {
		t.Errorf("Expected start time %v, got %v", params.StartTime, createdReservation.StartTime)
	}
	if !createdReservation.EndTime.Equal(params.EndTime) {
		t.Errorf("Expected end time %v, got %v", params.EndTime, createdReservation.EndTime)
	}
}

func TestReservationService_GetAllReservations(t *testing.T) {
	// 準備
	mockRepo := newMockReservationRepository()
	now := time.Now()
	expectedReservations := []*model.Reservation{
		model.NewReservation(now, now.Add(1*time.Hour)),
		model.NewReservation(now.Add(2*time.Hour), now.Add(3*time.Hour)),
	}

	mockRepo.findAllFunc = func() ([]*model.Reservation, error) {
		return expectedReservations, nil
	}

	service := NewReservationService(mockRepo)

	// 実行
	reservations, err := service.GetAllReservations()

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(reservations) != len(expectedReservations) {
		t.Errorf("Expected %d reservations, got %d", len(expectedReservations), len(reservations))
	}
	for i, r := range reservations {
		if r.ID != expectedReservations[i].ID {
			t.Errorf("Expected reservation ID %s, got %s", expectedReservations[i].ID, r.ID)
		}
	}
}

func TestReservationService_DeleteReservation(t *testing.T) {
	// 準備
	mockRepo := newMockReservationRepository()
	testID := "test-id"
	var deletedID string

	mockRepo.deleteFunc = func(id string) error {
		deletedID = id
		return nil
	}

	service := NewReservationService(mockRepo)

	// 実行
	err := service.DeleteReservation(testID)

	// 検証
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if deletedID != testID {
		t.Errorf("Expected ID %s to be deleted, got %s", testID, deletedID)
	}
}