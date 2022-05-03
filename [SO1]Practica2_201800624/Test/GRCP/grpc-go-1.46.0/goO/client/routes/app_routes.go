package routes

import (
	"client/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// routes


	app.Get("/", controllers.Inicio)

}