package utils

import (
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/go-playground/validator/v10"
)

var ValidatorPool = &sync.Pool{New: func() interface{} {
	return validator.New()
}}

func ShouldBind(ctx *fiber.Ctx, s interface{}) error {
	if err := ctx.BodyParser(s); err != nil {
		return err
	}
	v := ValidatorPool.Get().(*validator.Validate)
	return v.Struct(s)
}

func ShouldBindQuery(ctx *fiber.Ctx, s interface{}) error {
	if err := ctx.QueryParser(s); err != nil {
		return err
	}
	v := ValidatorPool.Get().(*validator.Validate)
	return v.Struct(s)
}
