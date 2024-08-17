package api

import (
	"context"
	"golang_biomtrik_login_fido/domain"
	"golang_biomtrik_login_fido/dto"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type userApi struct {
	userService domain.UserService
}

// fiber app ini untuk menyambungkan routenya, karena servicefiber utama itu adala app
func NewUserApi(app *fiber.App, userService domain.UserService) {
	ua := userApi{
		userService: userService,
	}
	app.Post("/user/register", ua.register)
}

func (ua *userApi) register(c *fiber.Ctx) error {
	// kita bikin context agar lebih mudah menangani erronya
	ca, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel() //

	var req dto.UserRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err := ua.userService.Register(ca, req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.Response[string]{
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(dto.Response[string]{
		Message: "User registered successfully",
	})
}
