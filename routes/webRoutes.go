package routes

import (
	"AplikasiWeb/controllers"

	"github.com/gofiber/fiber/v2"
)

func status(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "hello world !",
	})
}

func Setup(aplikasi *fiber.App) {
	//routes untuk data barang

	aplikasi.Get("/", status)
	aplikasi.Get("/DataBarang/Tampil", controllers.GetDataBarang)
	aplikasi.Get("/DataBarang/Tambah", controllers.ViewaddDataBarang)
	aplikasi.Get("/DataBarang/Tampil/:id", controllers.GetDataById)
	aplikasi.Post("/Tambah", controllers.AddDataBarang)
	aplikasi.Delete("/DataBarang/Delete/:id", controllers.DeleteDataBarang)
	aplikasi.Put("/DataBarang/Edit/:id", controllers.UpdateDataBarang)
}
