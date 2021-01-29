package handlers

import (
	"net/http"

	"github.com/asaskevich/govalidator"

	response "github.com/chi07/go-comment-service/internal/http/reponse"

	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

type VoteHandler struct {
	voteService VoteService
}

func NewVoteHandler(createCommentService VoteService) *VoteHandler {
	return &VoteHandler{voteService: createCommentService}
}

func (h *VoteHandler) Vote() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var requestBody models.Vote
		err := ctx.BodyParser(&requestBody)
		if err != nil {
			return response.Error(ctx, http.StatusBadRequest, http.StatusBadRequest, "cannot decode the body of vote request")
		}

		_, err = govalidator.ValidateStruct(&requestBody)
		if err != nil {
			return response.Error(ctx, http.StatusBadRequest, http.StatusBadRequest, err.Error())
		}

		result, err := h.voteService.VoteComment(ctx, &requestBody)

		return response.Success(ctx, http.StatusOK, result)
	}
}
