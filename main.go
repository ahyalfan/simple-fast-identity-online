package main

import (
	"golang_biomtrik_login_fido/internal/api"
	"golang_biomtrik_login_fido/internal/config"
	"golang_biomtrik_login_fido/internal/connection"
	"golang_biomtrik_login_fido/internal/repository"
	"golang_biomtrik_login_fido/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()

	conn := connection.GetDatabase(&cnf.Databases)

	userRepository := repository.NewUser(conn)
	challengeRepository := repository.NewChallenge(conn)

	userService := service.NewUserService(userRepository)
	challengeService := service.NewChallengeService(challengeRepository, userRepository)

	app := fiber.New()

	api.NewUserApi(app, userService)
	api.NewChallengeApi(app, challengeService)

	err := app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
	if err != nil {
		panic(err)
	}
}
