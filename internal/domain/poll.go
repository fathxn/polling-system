package domain

import (
	"context"
	"time"
)

type Option struct{}

type Poll struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Question  string    `json:"question"`
	Options   []Option  `json:"options"`
	CreatedAt time.Time `json:"created_at"`
	EndedAt   time.Time `json:"ended_at"`
}

type PollRepository interface {
	Create(ctx context.Context, poll *Poll) error
	GetByID(ctx context.Context, ID uint) (*Poll, error)
	List(ctx context.Context) ([]*Poll, error)
	Update(ctx context.Context, poll *Poll) error
	Delete(ctx context.Context, ID uint) error
}

type PollUsecase interface {
	CreatePoll(ctx context.Context, poll *Poll) error
	GetPoll(ctx context.Context, id uint) (*Poll, error)
	ListPolls(ctx context.Context) ([]*Poll, error)
	UpdatePoll(ctx context.Context, poll *Poll) error
	DeletePoll(ctx context.Context, id uint) error
}
