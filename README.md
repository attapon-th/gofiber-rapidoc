# gofiber-rapidoc


## Example

```go
package main

import (
	rapidoc "github.com/attapon-th/gofiber-rapidoc"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	
    // Endpoint: rapidoc 
	app.Get("/rapidoc/*", rapidoc.New(rapidoc.Config{
		Title:      "Pet Storage",
		HeaderText: "API Pet",
		SpecURL:    "/rapidoc/openapi.yaml", // get file: `./openapi.yaml` type: openapi v3
		LogoURL:    "https://redocly.github.io/redoc/petstore-logo.png",
	}))

	app.Listen("127.0.0.1:8888")

}

```