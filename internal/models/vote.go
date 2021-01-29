package models

import "time"

type Vote struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Type      string    `json:"type" bson:"type" valid:"required,in(like,dislike)"`
	CommentID string    `json:"commentID" bson:"comment_id" valid:"required"`
	UserID    string    `json:"userID" bson:"user_id"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updated_at"`
}

var (
	VoteTypeLike    = "like"
	VoteTypeDisLike = "dislike"
)
