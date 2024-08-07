package object

import "time"

type Status struct {
	ID        int       `json:"id,omitempty"`
	AccountID int       `json:"account_id,omitempty" db:"account_id"`
	URL       *string   `json:"url,omitempty" db:"url"`
	Content   string    `json:"status"`
	CreateAt  time.Time `json:"create_at,omitempty" db:"create_at"`
}

type StatusDetail struct {
	ID       int `json:"id"`
	Account  Account
	Content  string    `json:"status"`
	CreateAt time.Time `json:"create_at,omitempty"`
	// MediaAttachments Attachment
}

type StatusRequest struct {
	Status string `json:"string,omitempty"`
	Medias []Media
}

func NewStatus(content string) *Status {
	return &Status{
		Content:  content,
		CreateAt: time.Now(),
	}
}
