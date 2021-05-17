package routes

import (
	"github.com/WilliamCSA04/go-first-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
  route.Get("", controllers.GetTodos)
}

