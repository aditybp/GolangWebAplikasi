package database

import (
	"AplikasiWeb/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_username = "root"
const DB_password = ""
const DB_name = "belajar_golang"
const DB_host = "127.0.0.1"
const DB_port = "3306"

var DB *gorm.DB

func ConnectDB() {
	dsn := DB_username + ":" + DB_password + "@tcp" + "(" + DB_host + ":" + DB_port + ")/" + DB_name + "?" + "parseTime=true&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("tidak bisa koneksi database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.DataBarang{})

}
