package main

import (
	rapidoc "github.com/attapon-th/gofiber-rapidoc"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// app.Get("/rapidoc/*", rapidoc.New())

	app.Get("/rapidoc/*", rapidoc.New(rapidoc.Config{
		Title:      "Pet Storage",
		HeaderText: "API Pet",
		SpecURL:    "/rapidoc/openapi.yaml",
		LogoURL:    "https://redocly.github.io/redoc/petstore-logo.png",
	}))

	app.Listen("127.0.0.1:8888")

}
