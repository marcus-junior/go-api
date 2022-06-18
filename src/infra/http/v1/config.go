package apiV1

import "github.com/gofiber/fiber/v2"

func Configure(f *fiber.App) {
	api := f.Group("api/v1")

	empresa := api.Group("/empresa")
	empresa.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("/")
	})
	empresa.Get("/:idEmpresa", func(c *fiber.Ctx) error {
		return c.SendString("/:idEmpresa")
	})
	empresa.Get("/:idEmpresa/clientes", func(c *fiber.Ctx) error {
		return c.SendString("/:idEmpresa/clientes")
	})
}
