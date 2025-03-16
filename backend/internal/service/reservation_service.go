package service

import (
	"time"

	"github.com/naofumi-fujii/489-yoyaku/backend/internal/model"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/repository"
)

type ReservationService struct {
	repo repository.ReservationRepository
}

func NewReservationService(repo repository.ReservationRepository) *ReservationService {
	return &ReservationService{repo: repo}
}

type CreateReservationParams struct {
	StartTime time.Time `json:"startTime" validate:"required"`
	EndTime   time.Time `json:"endTime" validate:"required,gtfield=StartTime"`
}

func (s *ReservationService) CreateReservation(params CreateReservationParams) (*model.Reservation, error) {
	reservation := model.NewReservation(params.StartTime, params.EndTime)
	err := s.repo.Create(reservation)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (s *ReservationService) GetAllReservations() ([]*model.Reservation, error) {
	return s.repo.FindAll()
}

func (s *ReservationService) DeleteReservation(id string) error {
	return s.repo.Delete(id)
}
