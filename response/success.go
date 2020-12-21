package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type SuccessInstance struct {
	ctx        *fiber.Ctx
	data       *fiber.Map
	statusCode int
}

func (si SuccessInstance) WithCode(statusCode int) *SuccessInstance {
	si.statusCode = statusCode
	return &si
}

func (si SuccessInstance) WithData(data fiber.Map) *SuccessInstance {
	si.data = &data
	return &si
}

func (si *SuccessInstance) End() error {
	response := fiber.Map{"success": true}
	if si.data != nil {
		response["data"] = *si.data
	}
	return si.ctx.Status(si.statusCode).JSON(response)
}

func Success(ctx *fiber.Ctx) *SuccessInstance {
	return &SuccessInstance{ctx: ctx, statusCode: http.StatusOK}
}
