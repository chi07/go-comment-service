package models

type Comment struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	Type      string `json:"type" bson:"type"`
	Content   string `json:"content" bson:"content"`
	ArticleID int64  `json:"articleID" bson:"article_id"`
	ParentID  string `json:"parentID" bson:"parent_id"`
}

var (
	CommentTypePost = "post"
)
