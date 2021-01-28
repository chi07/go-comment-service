package models

import "time"

type Vote struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Type      string    `json:"type" bson:"type"`
	CommentID string    `json:"commentID" bson:"comment_id"`
	UserID    string    `json:"userID" bson:"user_id"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

var (
	VoteTypeLike    = "like"
	VoteTypeDisLike = "dislike"
)
