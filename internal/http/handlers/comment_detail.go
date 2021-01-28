package handlers

import (
	"net/http"

	response "github.com/chi07/go-comment-service/internal/http/reponse"

	"github.com/gofiber/fiber/v2"
)

type GetCommentHandler struct {
	updateCommentService GetCommentService
}

func NewGetCommentHandler(updateCommentService GetCommentService) *GetCommentHandler {
	return &GetCommentHandler{updateCommentService: updateCommentService}
}

func (h *GetCommentHandler) GetComment() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		comment, err := h.updateCommentService.GetComment(ctx, id)
		if err != nil {
			return response.Error(ctx, http.StatusInternalServerError, http.StatusInternalServerError, err.Error())
		}

		return response.Success(ctx, http.StatusOK, comment)
	}
}
