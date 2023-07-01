package middleware

import (
	"authentication/domain"
	"authentication/internal/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(userService domain.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := strings.ReplaceAll(ctx.Get("Authorization"), "Bearer ", "")
		if token == "" {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		user, err := userService.ValidateToken(ctx.Context(), token)
		if err != nil {
			return ctx.SendStatus(util.GetHttpStatusCode(err))
		}
		ctx.Locals("x-user", user)
		return ctx.Next()

	}
}
