package services

import (
	"time"

	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type UpdateVoteService struct {
	voteRepo VoteRepo
}

func NewUpdateVoteService(
	commentRepo VoteRepo) *UpdateVoteService {
	return &UpdateVoteService{
		voteRepo: commentRepo,
	}
}

func (s *UpdateVoteService) UpdateVote(ctx *fiber.Ctx, id string, v *models.Vote) error {
	v.UpdatedAt = time.Now()
	err := s.voteRepo.Update(ctx, id, v)
	return errors.Wrap(err, "cannot update vote")
}
