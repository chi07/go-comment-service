package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Error *ErrorResponse `json:"error,omitempty"`
	Data  interface{}    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

func Error(ctx *fiber.Ctx, httpCode int, errCode int, message string) error {
	ctx.Status(httpCode)
	return ctx.JSON(&fiber.Map{
		"success": false,
		"error": &ErrorResponse{
			Code:    errCode,
			Message: message,
		},
	})
}

func Success(ctx *fiber.Ctx, code int, data interface{}) error {
	ctx.Status(code)
	return ctx.JSON(&fiber.Map{
		"data": data,
	})
}
