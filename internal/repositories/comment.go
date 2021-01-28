package repositories

import (
	"fmt"

	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Comment struct {
	db *mongo.Database
}

func NewComment(db *mongo.Database) *Comment {
	return &Comment{db: db}
}

func (repo *Comment) GetAll(ctx *fiber.Ctx) ([]*models.Comment, error) {
	query := bson.D{{}}
	cursor, err := repo.db.Collection("comments").Find(ctx.Context(), query)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get comments from database")
	}

	var comments = make([]*models.Comment, 0)

	if err := cursor.All(ctx.Context(), &comments); err != nil {
		return nil, errors.Wrap(err, "cannot decode comments from database")
	}

	return comments, nil
}

func (repo *Comment) Get(ctx *fiber.Ctx, id string) (*models.Comment, error) {
	commentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("comment id = %s is not valid", id))
	}
	filter := bson.D{{Key: "_id", Value: commentID}}
	comment := &models.Comment{}

	err = repo.db.Collection("comments").FindOne(ctx.Context(), filter).Decode(comment)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get comments from database")
	}

	return comment, nil
}

func (repo *Comment) Create(ctx *fiber.Ctx, c *models.Comment) (interface{}, error) {
	collection := repo.db.Collection("comments")
	insertionResult, err := collection.InsertOne(ctx.Context(), c)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get insert comment into database")
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(ctx.Context(), filter)

	createdComment := &models.Comment{}
	err = createdRecord.Decode(createdComment)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get comment from database")
	}

	return createdComment, nil
}

func (repo *Comment) Update(ctx *fiber.Ctx, id string, c *models.Comment) error {
	commentID, err := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: commentID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "content", Value: c.Content},
				{Key: "parentID", Value: c.ParentID},
				{Key: "postID", Value: c.PostID},
			},
		},
	}

	err = repo.db.Collection("comments").FindOneAndUpdate(ctx.Context(), query, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.Wrap(err, "cannot match any documents in the collection")
		}
		return errors.Wrap(err, "internal server error from updating comment into DB")
	}

	return nil
}

func (repo *Comment) Delete(ctx *fiber.Ctx, id string) error {
	commentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "comment id is not valid")
	}
	query := bson.D{{Key: "_id", Value: commentID}}
	result, err := repo.db.Collection("comments").DeleteOne(ctx.Context(), &query)

	if err != nil {
		return errors.Wrap(err, "cannot delete comment from database")
	}

	if result.DeletedCount < 1 {
		return errors.Wrap(err, "cannot find comment with id: "+id+" in the database")
	}

	return nil
}
