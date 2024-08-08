package object

import "time"

type Timeline struct {
	ID       int64 `json:"id,omitempty"`
	Account  Account
	Content  string    `json:"status,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
}
