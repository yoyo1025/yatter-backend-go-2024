package object

import "time"

type Status struct {
	ID        int       `json:"id,omitempty"`
	AccountID int       `json:"account_id,omitempty" db:"account_id"`
	URL       *string   `json:"url,omitempty" db:"url"`
	Content   string    `json:"status"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

type StatusDetail struct {
	ID               int `json:"id"`
	Account          Account
	Content          string    `json:"status"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	MediaAttachments Attachment
}

func NewStatus(content string) *Status {
	return &Status{
		Content:   content,
		CreatedAt: time.Now(),
	}
}
