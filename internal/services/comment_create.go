package services

import (
	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type CreateCommentService struct {
	commentRepo CreateCommentRepo
}

func NewCreateCommentService(
	commentRepo CreateCommentRepo) *CreateCommentService {
	return &CreateCommentService{
		commentRepo: commentRepo,
	}
}

func (s *CreateCommentService) CreateComment(ctx *fiber.Ctx, c *models.Comment) (interface{}, error) {
	if c.Type == "" {
		c.Type = models.CommentTypePost
	}
	comment, err := s.commentRepo.Create(ctx, c)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create comment")
	}
	return comment, nil
}
