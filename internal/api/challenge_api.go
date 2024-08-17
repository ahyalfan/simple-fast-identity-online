package api

import (
	"context"
	"golang_biomtrik_login_fido/domain"
	"golang_biomtrik_login_fido/dto"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type challengeApi struct {
	challengeService domain.ChallengeService
}

func NewChallengeApi(app *fiber.App, challengeService domain.ChallengeService) {
	ca := challengeApi{
		challengeService: challengeService,
	}
	app.Get("/challenge", ca.generate)
	app.Post("/challenge", ca.validate)
}

func (c *challengeApi) generate(ctx *fiber.Ctx) error {
	// kita bikin context agar lebih mudah menangani erronya
	ca, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel() //

	challenge, err := c.challengeService.Generate(ca)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response[string]{
			Message: err.Error(),
		})
	}
	return ctx.JSON(dto.Response[dto.ChallengeData]{
		Data: challenge,
	})
}

func (c *challengeApi) validate(ctx *fiber.Ctx) error {
	// kita bikin context agar lebih mudah menangani erronya
	ca, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel() //

	var req dto.ChallengeValidate
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	challenge, err := c.challengeService.Validate(ca, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response[string]{
			Message: err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(dto.Response[dto.UserData]{
		Data: challenge,
	})
}
