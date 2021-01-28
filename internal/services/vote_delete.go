package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type DeleteVoteService struct {
	voteRepo VoteRepo
}

func NewDeleteVoteService(
	voteRepo VoteRepo) *DeleteVoteService {
	return &DeleteVoteService{
		voteRepo: voteRepo,
	}
}

func (s *DeleteVoteService) DeleteVote(ctx *fiber.Ctx, id string) error {
	err := s.voteRepo.Delete(ctx, id)
	return errors.Wrap(err, "cannot delete the vote with id: "+id)
}
