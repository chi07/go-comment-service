package services

import (
	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

type VoteRepo interface {
	Create(ctx *fiber.Ctx, c *models.Vote) (*models.Vote, error)
	Update(ctx *fiber.Ctx, id string, c *models.Vote) error
	Delete(ctx *fiber.Ctx, id string) error
	Get(ctx *fiber.Ctx, commentID, userID string) (*models.Vote, error)
	CountVote(ctx *fiber.Ctx, commentID, voteType string) (int64, error)
}
