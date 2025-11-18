package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func currentTime(c *fiber.Ctx) error {
	currentTime := time.Now()
	return c.SendString(currentTime.String())
}

func main() {

	server := fiber.New()

	server.Get("/", func(c *fiber.Ctx) error {
		return  c.SendString("Hello, World!")
	})

	server.Get("/ruta1", currentTime)

	server.Listen(":3000")
	
}
