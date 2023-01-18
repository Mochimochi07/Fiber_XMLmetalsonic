package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"encoding/xml"
)

type metalSonic struct {
	XMLName xml.Name `xml:"metal-sonic"`
	ID      int      `xml:"id"`
	Name    string   `xml:"name"`
	Ability string   `xml:"ability"`
}

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	app.Get("/", func(c *fiber.Ctx) {
		metal := &metalSonic{ID: 1, Name: "Metal Sonic", Ability: "Super speed"}
		c.XML(metal)
	})

	app.Listen(3000)
}
