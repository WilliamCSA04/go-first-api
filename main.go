package main

import (
	"github.com/WilliamCSA04/go-first-api/routes"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
  app.Get("/", func(c *fiber.Ctx) error {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "success":  true,
        "message": "You are at the endpoint 😉",
    })
  })
  
  api := app.Group("/api")

  api.Get("", func(c *fiber.Ctx) error {
      return c.Status(fiber.StatusOK).JSON(fiber.Map{
          "success": true,
          "message": "You are at the api endpoint 😉",
      })
  })

  routes.TodoRoute(api.Group("/todos"))
}

func main() {
  app := fiber.New()

  setupRoutes(app)

  err := app.Listen(":8000")

  if err != nil {
      panic(err)
  }

}
