package services

import (
	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CommentRepo interface {
	Create(ctx *fiber.Ctx, c *models.Comment) (interface{}, error)
	Update(ctx *fiber.Ctx, id string, c *models.Comment) error
	Delete(ctx *fiber.Ctx, id string) error
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
