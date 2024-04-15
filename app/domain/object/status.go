package object

import (
	"time"
)

type Status struct {
	// The internal ID of the account
	ID int64 `json:"id,omitempty"`

	// The internal ID of the account
	AccountID int64 `db:"account_id"`

	// The content of the status
	Status string `json:"status,omitempty" db:"content"`

	// The time the account was created
	CreateAt time.Time `json:"create_at,omitempty" db:"create_at"`
}
