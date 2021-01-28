package models

import "time"

type Comment struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Type      string    `json:"type" bson:"type"`
	Content   string    `json:"content" bson:"content"`
	PostID    int64     `json:"postID" bson:"post_id"`
	ParentID  string    `json:"parentID" bson:"parent_id"`
	UserID    string    `json:"userID" bson:"user_id"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

var (
	CommentTypePost = "post"
)
