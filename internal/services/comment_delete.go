package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type DeleteCommentService struct {
	commentRepo DeleteCommentRepo
}

func NewDeleteCommentService(
	commentRepo DeleteCommentRepo) *DeleteCommentService {
	return &DeleteCommentService{
		commentRepo: commentRepo,
	}
}

func (s *DeleteCommentService) DeleteComment(ctx *fiber.Ctx, id string) error {
	err := s.commentRepo.Delete(ctx, id)
	if err != nil {
		return errors.Wrap(err, "cannot delete the comment")
	}
	return nil
}
