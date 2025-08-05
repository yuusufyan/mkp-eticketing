package main

import (
	"auth-rbac/src/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Logger middleware ala Air (warna dan detail waktu)
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip} - ${status} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	config.ConnectDB()

	app.Listen(":3000")
}
