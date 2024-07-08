package repository

import (
	"context"
	"database/sql"
	"polling-system/internal/domain"
)

type pollRepository struct {
	db *sql.DB
}

func NewPollRepository(db *sql.DB) domain.PollRepository {
	return &pollRepository{db: db}
}

// Create implements domain.PollRepository.
func (p *pollRepository) Create(ctx context.Context, poll *domain.Poll) error {
	panic("unimplemented")
}

// Delete implements domain.PollRepository.
func (p *pollRepository) Delete(ctx context.Context, ID uint) error {
	panic("unimplemented")
}

// GetByID implements domain.PollRepository.
func (p *pollRepository) GetByID(ctx context.Context, ID uint) (*domain.Poll, error) {
	panic("unimplemented")
}

// List implements domain.PollRepository.
func (p *pollRepository) List(ctx context.Context) ([]*domain.Poll, error) {
	panic("unimplemented")
}

// Update implements domain.PollRepository.
func (p *pollRepository) Update(ctx context.Context, poll *domain.Poll) error {
	panic("unimplemented")
}
