package handlers

import (
	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CreateCommentService interface {
	CreateComment(ctx *fiber.Ctx, c *models.Comment) (interface{}, error)
}

type UpdateCommentService interface {
	UpdateComment(ctx *fiber.Ctx, id string, c *models.Comment) error
}

type DeleteCommentService interface {
	DeleteComment(ctx *fiber.Ctx, id string) error
}

type GetCommentService interface {
	GetComment(ctx *fiber.Ctx, id string) (*models.Comment, error)
}
