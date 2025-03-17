package repository

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/naofumi-fujii/489-yoyaku/backend/internal/model"
)

type ReservationRepository interface {
	Create(reservation *model.Reservation) error
	FindAll() ([]*model.Reservation, error)
	FindByID(id string) (*model.Reservation, error)
	Delete(id string) error
}

// InMemoryReservationRepository - Keep for backward compatibility and testing
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

// MySQLReservationRepository - MySQL implementation
type MySQLReservationRepository struct {
	db *sql.DB
}

// NewMySQLReservationRepository creates a new MySQL repository
func NewMySQLReservationRepository(db *sql.DB) (*MySQLReservationRepository, error) {
	// Create reservations table if it doesn't exist
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS reservations (
			id VARCHAR(36) PRIMARY KEY,
			start_time DATETIME NOT NULL,
			end_time DATETIME NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create reservations table: %w", err)
	}

	return &MySQLReservationRepository{
		db: db,
	}, nil
}

// Create inserts a new reservation into the database
func (r *MySQLReservationRepository) Create(reservation *model.Reservation) error {
	_, err := r.db.Exec(
		"INSERT INTO reservations (id, start_time, end_time, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		reservation.ID,
		reservation.StartTime,
		reservation.EndTime,
		reservation.CreatedAt,
		reservation.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create reservation: %w", err)
	}
	return nil
}

// FindAll returns all reservations
func (r *MySQLReservationRepository) FindAll() ([]*model.Reservation, error) {
	rows, err := r.db.Query("SELECT id, start_time, end_time, created_at, updated_at FROM reservations")
	if err != nil {
		return nil, fmt.Errorf("failed to find reservations: %w", err)
	}
	defer rows.Close()

	var reservations []*model.Reservation
	for rows.Next() {
		var reservation model.Reservation
		var startTime, endTime, createdAt, updatedAt time.Time
		if err := rows.Scan(&reservation.ID, &startTime, &endTime, &createdAt, &updatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan reservation: %w", err)
		}
		reservation.StartTime = startTime
		reservation.EndTime = endTime
		reservation.CreatedAt = createdAt
		reservation.UpdatedAt = updatedAt
		reservations = append(reservations, &reservation)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return reservations, nil
}

// FindByID returns a reservation by ID
func (r *MySQLReservationRepository) FindByID(id string) (*model.Reservation, error) {
	var reservation model.Reservation
	var startTime, endTime, createdAt, updatedAt time.Time
	
	err := r.db.QueryRow(
		"SELECT id, start_time, end_time, created_at, updated_at FROM reservations WHERE id = ?", 
		id,
	).Scan(&reservation.ID, &startTime, &endTime, &createdAt, &updatedAt)
	
	if err == sql.ErrNoRows {
		return nil, nil
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to find reservation: %w", err)
	}
	
	reservation.StartTime = startTime
	reservation.EndTime = endTime
	reservation.CreatedAt = createdAt
	reservation.UpdatedAt = updatedAt
	
	return &reservation, nil
}

// Delete removes a reservation by ID
func (r *MySQLReservationRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM reservations WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete reservation: %w", err)
	}
	return nil
}
