package main

import (
	"AplikasiWeb/database"
	"AplikasiWeb/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	database.ConnectDB()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	/**
	app.Get("/contoh_halaman", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "hello world !",
		})
	})*/

	routes.Setup(app)

	app.Listen(":8000")
}
