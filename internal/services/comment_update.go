package services

import (
	"time"

	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type UpdateCommentService struct {
	commentRepo UpdateCommentRepo
}

func NewUpdateCommentService(
	commentRepo UpdateCommentRepo) *UpdateCommentService {
	return &UpdateCommentService{
		commentRepo: commentRepo,
	}
}

func (s *UpdateCommentService) UpdateComment(ctx *fiber.Ctx, id string, c *models.Comment) error {
	c.UpdatedAt = time.Now()
	err := s.commentRepo.Update(ctx, id, c)
	if err != nil {
		return errors.Wrap(err, "cannot update comment")
	}
	return nil
}
