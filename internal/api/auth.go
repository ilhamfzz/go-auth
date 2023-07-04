package api

import (
	"authentication/domain"
	"authentication/dto"
	"authentication/internal/util"

	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	userService domain.UserService
}

func NewAuth(app *fiber.App, userService domain.UserService, authMiddleware fiber.Handler) {
	api := authApi{
		userService: userService,
	}

	app.Post("user/register", api.Register)
	app.Post("user/validate-otp", api.ValidateOTP)
	app.Post("token/generate", api.GenerateToken)
	app.Get("token/validate", authMiddleware, api.ValidateToken)
}

func (a authApi) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	user, err := a.userService.Register(ctx.Context(), req)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatusCode(err))
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func (a authApi) GenerateToken(ctx *fiber.Ctx) error {
	var req dto.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	token, err := a.userService.Authenticate(ctx.Context(), req)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatusCode(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(token)
}

func (a authApi) ValidateToken(ctx *fiber.Ctx) error {
	user := ctx.Locals("x-user")
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (a authApi) ValidateOTP(ctx *fiber.Ctx) error {
	var req dto.ValidateOTPReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err := a.userService.ValidateOTP(ctx.Context(), req)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatusCode(err))
	}

	return ctx.SendStatus(fiber.StatusOK)
}
