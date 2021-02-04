package object

import (
	"time"

	"github.com/go-gorp/gorp/v3"
)

type (
	StatusID = int64

	// Status status
	Status struct {
		// The ID of the status
		ID StatusID `json:"id,omitempty"`

		// account id
		AccountID AccountID `json:"-" db:"account_id"`

		// account
		Account *Account `json:"account,omitempty" db:"-"`

		// Body of the status; this will contain HTML (remote HTML already sanitized)
		Content string `json:"content,omitempty"`

		// URL to the status page (can be remote)
		URL string `json:"url,omitempty"`

		// The time the status was created
		// Format: date-time
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}

	FindStatusCondition struct {
		Limit   int   `json:"limit"`
		SinceID int64 `json:"since_id"`
		MaxID   int64 `json:"max_id"`
	}
)

func (st *Status) PreInsert(s gorp.SqlExecutor) error {
	st.CreateAt = DateTime{time.Now()}

	if st.Account != nil && st.Account.ID > 0 {
		st.AccountID = st.Account.ID
	}

	return nil
}

func (st *Status) PostGet(s gorp.SqlExecutor) error {
	if st.AccountID > 0 {
		// TODO: select join
		st.Account = &Account{
			ID: st.AccountID,
		}
	}

	return nil
}
