package services

import (
	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

type BookmarkRepo interface {
	Create(ctx *fiber.Ctx, c *models.Bookmark) (*models.Bookmark, error)
	Delete(ctx *fiber.Ctx, id string) error
	Get(ctx *fiber.Ctx, postID, userID string) (*models.Bookmark, error)
}
