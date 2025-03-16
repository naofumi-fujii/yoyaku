package repository

import (
	"sync"

	"github.com/naofumi-fujii/489-yoyaku/backend/internal/model"
)

type ReservationRepository interface {
	Create(reservation *model.Reservation) error
	FindAll() ([]*model.Reservation, error)
	FindByID(id string) (*model.Reservation, error)
	Delete(id string) error
}

type InMemoryReservationRepository struct {
	reservations map[string]*model.Reservation
	mutex        sync.RWMutex
}

func NewInMemoryReservationRepository() *InMemoryReservationRepository {
	return &InMemoryReservationRepository{
		reservations: make(map[string]*model.Reservation),
	}
}

func (r *InMemoryReservationRepository) Create(reservation *model.Reservation) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.reservations[reservation.ID] = reservation
	return nil
}

func (r *InMemoryReservationRepository) FindAll() ([]*model.Reservation, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	reservations := make([]*model.Reservation, 0, len(r.reservations))
	for _, reservation := range r.reservations {
		reservations = append(reservations, reservation)
	}

	return reservations, nil
}

func (r *InMemoryReservationRepository) FindByID(id string) (*model.Reservation, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	reservation, ok := r.reservations[id]
	if !ok {
		return nil, nil
	}

	return reservation, nil
}

func (r *InMemoryReservationRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	delete(r.reservations, id)
	return nil
}
