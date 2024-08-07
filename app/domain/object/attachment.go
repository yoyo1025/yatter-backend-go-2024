package object

type Attachment struct {
	ID          int64  `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Url         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}
