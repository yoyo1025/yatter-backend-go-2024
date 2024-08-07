package object

type Relationship struct {
	FollowerID int64 `json:"follower_id,omitempty"`
	FolloweeID int64 `json:"followee_id,omitempty"`
}
