package models

import "time"

type Vote struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Type      string    `json:"type" bson:"type"`
	CommentID string    `json:"commentID" bson:"comment_id"`
	UserID    string    `json:"userID" bson:"user_id"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
}

var (
	VoteTypeLike    = "like"
	VoteTypeDisLike = "dislike"
)
