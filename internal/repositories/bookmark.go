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

type Bookmark struct {
	db *mongo.Database
}

func NewBookmark(db *mongo.Database) *Bookmark {
	return &Bookmark{db: db}
}

const bookmarkTable = "bookmarks"

func (repo *Bookmark) Get(ctx *fiber.Ctx, postID, userID string) (*models.Bookmark, error) {
	var err error
	filter := bson.D{primitive.E{Key: "post_id", Value: postID}, primitive.E{Key: "user_id", Value: userID}}
	var vote *models.Bookmark

	err = repo.db.Collection(bookmarkTable).FindOne(ctx.Context(), filter).Decode(&vote)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, errors.Wrap(err, fmt.Sprintf("cannot get vote from database with postID=%s and userID=%s", postID, userID))
	}

	return vote, nil
}

func (repo *Bookmark) Create(ctx *fiber.Ctx, c *models.Bookmark) (*models.Bookmark, error) {
	collection := repo.db.Collection(bookmarkTable)
	insertionResult, err := collection.InsertOne(ctx.Context(), c)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get insert bookmark post into database")
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(ctx.Context(), filter)

	createdBookmark := &models.Bookmark{}
	err = createdRecord.Decode(createdBookmark)
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode the bookmark post from database")
	}

	return createdBookmark, nil
}

func (repo *Bookmark) Delete(ctx *fiber.Ctx, id string) error {
	voteID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "vote id is not valid")
	}
	query := bson.D{{Key: "_id", Value: voteID}}
	result, err := repo.db.Collection(bookmarkTable).DeleteOne(ctx.Context(), &query)

	if err != nil {
		return errors.Wrap(err, "cannot delete bookmark from database for id: "+id)
	}

	if result.DeletedCount < 1 {
		return errors.Wrap(err, "cannot find bookmark with id: "+id+" in the database")
	}

	return nil
}
