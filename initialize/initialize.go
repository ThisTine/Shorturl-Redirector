package initialize

import (
	"os"

	routes "github.com/ThisTine/shorturl-redirector/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Initialize(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("APP_ORIGIN"),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://service.thistine.com/")
	})

	app.Get("/:pathid", routes.Redirect)
}
