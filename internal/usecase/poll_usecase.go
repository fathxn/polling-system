package usecase

import (
	"context"
	"polling-system/internal/domain"
)

type pollUsecase struct {
	pollRepository domain.PollRepository
}

func NewPollUsecase(pollRepository domain.PollRepository) domain.PollUsecase {
	return &pollUsecase{pollRepository: pollRepository}
}

// CreatePoll implements domain.PollUsecase.
func (p *pollUsecase) CreatePoll(ctx context.Context, poll *domain.Poll) error {
	panic("unimplemented")
}

// DeletePoll implements domain.PollUsecase.
func (p *pollUsecase) DeletePoll(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// GetPoll implements domain.PollUsecase.
func (p *pollUsecase) GetPoll(ctx context.Context, id uint) (*domain.Poll, error) {
	panic("unimplemented")
}

// ListPolls implements domain.PollUsecase.
func (p *pollUsecase) ListPolls(ctx context.Context) ([]*domain.Poll, error) {
	panic("unimplemented")
}

// UpdatePoll implements domain.PollUsecase.
func (p *pollUsecase) UpdatePoll(ctx context.Context, poll *domain.Poll) error {
	panic("unimplemented")
}
