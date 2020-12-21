package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorInstance struct {
	ctx        *fiber.Ctx
	data       *fiber.Map
	statusCode int
	message    *string
}

func (ei ErrorInstance) WithCode(statusCode int) *ErrorInstance {
	ei.statusCode = statusCode
	return &ei
}

func (ei ErrorInstance) WithMessage(message string) *ErrorInstance {
	ei.message = &message
	return &ei
}

func (ei ErrorInstance) WithData(data fiber.Map) *ErrorInstance {
	ei.data = &data
	return &ei
}

func (ei *ErrorInstance) End() error {
	err := fiber.Map{"code": ei.statusCode}
	if ei.message != nil {
		err["message"] = *ei.message
	}

	response := fiber.Map{"error": err}
	if ei.data != nil {
		response["data"] = *ei.data
	}
	return ei.ctx.Status(ei.statusCode).JSON(response)
}

func Error(ctx *fiber.Ctx) *ErrorInstance {
	return &ErrorInstance{ctx: ctx, statusCode: http.StatusBadRequest}
}
