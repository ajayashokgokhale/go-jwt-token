package main

import (
	"jwt-token/pkg/dbx"
	"jwt-token/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	dbx.Init()

	app := fiber.New()
	request.RegisterRoutes(app)

	port := ":8080"
	log.Printf("Server running on http://localhost%s", port)
	log.Fatal(app.Listen(port))
}
