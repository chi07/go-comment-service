package handlers

import (
	"net/http"

	response "github.com/chi07/go-comment-service/internal/http/reponse"

	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateCommentHandler struct {
	updateCommentService UpdateCommentService
}

func NewUpdateCommentHandler(updateCommentService UpdateCommentService) *UpdateCommentHandler {
	return &UpdateCommentHandler{updateCommentService: updateCommentService}
}

func (h *UpdateCommentHandler) UpdateComment() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var requestBody models.Comment
		err := ctx.BodyParser(&requestBody)
		if err != nil {
			return response.Error(ctx, http.StatusBadRequest, http.StatusBadRequest, "cannot decode the body request")
		}
		err = h.updateCommentService.UpdateComment(ctx, id, &requestBody)
		if err != nil {
			return response.Error(ctx, http.StatusInternalServerError, http.StatusInternalServerError, err.Error())
		}

		return response.Success(ctx, http.StatusOK, nil)
	}
}
