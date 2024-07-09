package domain

import "time"

type Vote struct {
	ID       uint      `json:"id"`
	PollID   uint      `json:"poll_id"`
	OptionID uint      `json:"option_id"`
	VoterID  string    `json:"voter_id"`
	VotedAt  time.Time `json:"voted_at"`
}
