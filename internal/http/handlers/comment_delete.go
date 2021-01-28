package handlers

import (
	"net/http"

	response "github.com/chi07/go-comment-service/internal/http/reponse"

	"github.com/gofiber/fiber/v2"
)

type DeleteCommentHandler struct {
	deleteCommentService DeleteCommentService
}

func NewDeleteCommentHandler(deleteCommentService DeleteCommentService) *DeleteCommentHandler {
	return &DeleteCommentHandler{deleteCommentService: deleteCommentService}
}

func (h *DeleteCommentHandler) DeleteComment() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		err := h.deleteCommentService.DeleteComment(ctx, id)
		if err != nil {
			return response.Error(ctx, http.StatusInternalServerError, http.StatusInternalServerError, err.Error())
		}

		return response.Success(ctx, http.StatusOK, nil)
	}
}
