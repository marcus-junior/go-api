package main

import (
	"github.com/gofiber/fiber/v2"
	apiV1 "github.com/marcus-junior/go-api/src/infra/http/v1"
)

func main() {

	app := fiber.New()
	apiV1.Configure(app)

	app.Listen(":3000")
}
