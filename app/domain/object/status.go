package object

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
