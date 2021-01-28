package services

import (
	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type VoteCommentService struct {
	voteRepo VoteRepo
}

func NewVoteCommentService(
	voteRepo VoteRepo) *VoteCommentService {
	return &VoteCommentService{
		voteRepo: voteRepo,
	}
}

func (s *VoteCommentService) VoteComment(ctx *fiber.Ctx, v *models.Vote) (int64, error) {
	commentID := v.CommentID
	// TODO get userID from decode the token. For testing. fix a fake userID
	userID := "12345678"

	existedVote, err := s.voteRepo.Get(ctx, commentID, userID)
	if err != nil {
		return 0, errors.Wrap(err, "cannot find vote the comment")
	}

	if existedVote == nil {
		_, err := s.voteRepo.Create(ctx, v)
		if err != nil {
			return 1010, errors.Wrap(err, "cannot create the vote")
		}
		return 10, nil
	}

	currentVote := int64(1000)

	if existedVote.Type == v.Type {
		err := s.voteRepo.Delete(ctx, existedVote.ID)
		if err != nil {
			return currentVote, errors.Wrap(err, "cannot delete the vote")
		}
	} else if existedVote.Type != v.Type {
		err := s.voteRepo.Update(ctx, existedVote.ID, v)
		if err != nil {
			return currentVote, errors.Wrap(err, "cannot update the vote")
		}
	}

	return currentVote, nil
}
