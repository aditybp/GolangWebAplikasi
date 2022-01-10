package controllers

import (
	"AplikasiWeb/database"
	"AplikasiWeb/models"

	"github.com/gofiber/fiber/v2"
)

func GetDataBarang(c *fiber.Ctx) error {
	var databarang []models.DataBarang
	database.DB.Find(&databarang)

	if len(databarang) == 0 {
		return c.Status(404).JSON(fiber.Map{"Status": "eror", "data": nil})
	}

	return c.Render("getalldata", fiber.Map{		
		"databarang": databarang,
	})

	//return c.JSON(databarang)
}

func ViewaddDataBarang(c *fiber.Ctx) error {
	return c.Render("inputdata", fiber.Map{
		"Title": "halaman edit",
	})
}

func AddDataBarang(c *fiber.Ctx) error {
	data_barang := new(models.DataBarang)

	if err := c.BodyParser(data_barang); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	input := models.DataBarang{
		Nama_Barang:   c.FormValue("nama_barang"),
		Kode_Barang:   c.FormValue("kode_barang"),
		Jumlah_Barang: c.FormValue("jumlah_barang"),
	}

	database.DB.Create(&input)
	//return c.JSON(&data_barang)
	return c.Redirect("/DataBarang/Tampil")
}

func UpdateDataBarang(c *fiber.Ctx) error {
	id := c.Params("id")
	data_barang := new(models.DataBarang)
	database.DB.First(&data_barang, id)
	if data_barang.Nama_Barang == "" {
		return c.Status(500).SendString("data tidak tersedia")
	}
	if err := c.BodyParser(&data_barang); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Save(&data_barang)
	return c.JSON(&data_barang)
}

func DeleteDataBarang(c *fiber.Ctx) error {
	id := c.Params("id")
	var data_barang models.DataBarang
	database.DB.First(&data_barang, id)

	if data_barang.Nama_Barang == "" {
		c.Status(500).SendString("tidak ada data")

	}
	database.DB.Delete(&data_barang)
	return c.SendString("data terhapus")
}

func GetDataById(c *fiber.Ctx) error {
	id := c.Params("id")
	var data_barang models.DataBarang
	database.DB.Find(&data_barang, id)
	//return c.JSON(&data_barang)
	return c.Render("index", fiber.Map{
		"judul": data_barang.Id,
		"Title": data_barang.Nama_Barang,
	})
}
