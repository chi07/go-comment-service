package services

import (
	"time"

	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type CreateVoteService struct {
	voteRepo VoteRepo
}

func NewCreateVoteService(
	voteRepo VoteRepo) *CreateVoteService {
	return &CreateVoteService{
		voteRepo: voteRepo,
	}
}

func (s *CreateVoteService) CreateVote(ctx *fiber.Ctx, v *models.Vote) (interface{}, error) {
	now := time.Now()
	v.CreatedAt = now
	v.UpdatedAt = now

	// TODO get userID from decode the token. For testing. fix a fake userID
	v.UserID = "12345678"
	if v.Type == "" {
		v.Type = models.VoteTypeLike
	}
	comment, err := s.voteRepo.Create(ctx, v)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create vote")
	}
	return comment, nil
}
