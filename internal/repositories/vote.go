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

type Vote struct {
	db *mongo.Database
}

func NewVote(db *mongo.Database) *Vote {
	return &Vote{db: db}
}

const voteTables = "votes"

func (repo *Vote) Get(ctx *fiber.Ctx, commentID, userID string) (*models.Vote, error) {
	filter := bson.D{{Key: "comment_id", Value: commentID}, {Key: "user_id", Value: userID}}
	var vote *models.Vote

	err := repo.db.Collection(voteTables).FindOne(ctx.Context(), filter).Decode(vote)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("cannot get vote from database with commentID=%s and userID=%s", commentID, userID))
	}

	return vote, nil
}

func (repo *Vote) Create(ctx *fiber.Ctx, c *models.Vote) (*models.Vote, error) {
	collection := repo.db.Collection(voteTables)
	insertionResult, err := collection.InsertOne(ctx.Context(), c)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get insert comment into database")
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(ctx.Context(), filter)

	var createdVote *models.Vote
	err = createdRecord.Decode(createdVote)
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode the vote comment from database")
	}

	return createdVote, nil
}

func (repo *Vote) Update(ctx *fiber.Ctx, id string, c *models.Vote) error {
	voteID, err := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: voteID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "type", Value: c.Type},
			},
		},
	}

	err = repo.db.Collection(voteTables).FindOneAndUpdate(ctx.Context(), query, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.Wrap(err, "cannot match any documents in the votes collection")
		}
		return errors.Wrap(err, "internal server error from updating vote into DB")
	}

	return nil
}

func (repo *Vote) Delete(ctx *fiber.Ctx, id string) error {
	voteID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "vote id is not valid")
	}
	query := bson.D{{Key: "_id", Value: voteID}}
	result, err := repo.db.Collection(voteTables).DeleteOne(ctx.Context(), &query)

	if err != nil {
		return errors.Wrap(err, "cannot delete vote from database for id: "+id)
	}

	if result.DeletedCount < 1 {
		return errors.Wrap(err, "cannot find vote with id: "+id+" in the database")
	}

	return nil
}
