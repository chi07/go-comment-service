package handlers

import (
	"net/http"

	response "github.com/chi07/go-comment-service/internal/http/reponse"

	"github.com/chi07/go-comment-service/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CreateCommentHandler struct {
	createCommentService CreateCommentService
}

func NewCreateCommentHandler(createCommentService CreateCommentService) *CreateCommentHandler {
	return &CreateCommentHandler{createCommentService: createCommentService}
}

func (h *CreateCommentHandler) CreateComment() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var requestBody models.Comment
		err := ctx.BodyParser(&requestBody)
		if err != nil {
			return response.Error(ctx, http.StatusBadRequest, http.StatusBadRequest, "cannot decode the body request")
		}
		result, err := h.createCommentService.CreateComment(ctx, &requestBody)

		return response.Success(ctx, http.StatusCreated, result)
	}
}
