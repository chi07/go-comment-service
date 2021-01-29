package models

import "time"

type Bookmark struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	PostID    string    `json:"postID" bson:"post_id" valid:"required"`
	UserID    string    `json:"userID" bson:"user_id"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"created_at"`
}
