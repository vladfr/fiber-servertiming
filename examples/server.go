package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberservertiming "github.com/vladfr/fiber-servertiming/v2"
)

func main() {
	app := fiber.New()
	app.Use(fiberservertiming.New(fiberservertiming.Config{
		AllowOrigins: "http://127.0.0.1:3000, http://127.0.0.1:3000/",
	}))

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./index.html")
	})

	app.Get("/server", func(c *fiber.Ctx) error {
		timing := fiberservertiming.FromContext(c)
		time.Sleep(12 * time.Millisecond)

		defer timing.NewMetric("backendcall").Start().Stop()
		time.Sleep(10 * time.Millisecond)

		return c.SendString(fmt.Sprintf("Hello from the server at %s", time.Now()))
	})

	port := "3000"
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
