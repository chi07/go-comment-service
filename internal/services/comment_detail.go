package services

import (
	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type GetCommentService struct {
	commentRepo GetCommentRepo
}

func NewGetCommentService(
	commentRepo GetCommentRepo) *GetCommentService {
	return &GetCommentService{
		commentRepo: commentRepo,
	}
}

func (s *GetCommentService) GetComment(ctx *fiber.Ctx, id string) (*models.Comment, error) {
	comment, err := s.commentRepo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get comment")
	}
	return comment, nil
}
