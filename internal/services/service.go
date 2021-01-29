package services

import (
	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CommentRepo interface {
	Create(ctx *fiber.Ctx, c *models.Comment) (interface{}, error)
	Update(ctx *fiber.Ctx, id string, c *models.Comment) error
	Delete(ctx *fiber.Ctx, id string) error
	Get(ctx *fiber.Ctx, id string) (*models.Comment, error)
}

type CreateCommentRepo interface {
	Create(ctx *fiber.Ctx, c *models.Comment) (interface{}, error)
}

type UpdateCommentRepo interface {
	Update(ctx *fiber.Ctx, id string, c *models.Comment) error
}

type DeleteCommentRepo interface {
	Delete(ctx *fiber.Ctx, id string) error
}

type GetCommentRepo interface {
	Get(ctx *fiber.Ctx, id string) (*models.Comment, error)
}

type VoteRepo interface {
	Create(ctx *fiber.Ctx, c *models.Vote) (*models.Vote, error)
	Update(ctx *fiber.Ctx, id string, c *models.Vote) error
	Delete(ctx *fiber.Ctx, id string) error
	Get(ctx *fiber.Ctx, commentID, userID string) (*models.Vote, error)
	CountVote(ctx *fiber.Ctx, commentID, voteType string) (int64, error)
}
