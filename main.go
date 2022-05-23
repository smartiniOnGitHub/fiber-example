package main

import (
	// "encoding/json"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

const (
	bind_address string = "0.0.0.0" // empty means all available interfaces, set '0.0.0.0' here just for explitic value
	bind_port    string = "8000"
	address      string = bind_address + ":" + bind_port
)

func main() {
	app := fiber.New()

	// TODO: update home page handler to return some static content ... wip
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-type", "application/json")
		return c.SendString(`{"msg":"Hello from Fiber"}`)
	})

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		// j, err := json.Marshal(items)
		// if err != nil { return err }
		// c.Set("Content-type", "application/json; charset=utf-8")
		c.Set("Content-type", "application/json") // should be enough
		// return c.Send(j)
		return c.SendString(`{"Hello":"World"}`)
	})

	app.Get("/api/info", func(c *fiber.Ctx) error {
		c.Set("Content-type", "application/json")
		return c.SendString(`{"msg":"Hello from Fiber"}`)
	})

	// TODO: add other routes (maybe in its own source), handlers, etc ... wip

	fmt.Printf("Server listening on 'http://%s' ...\n", address)
	app.Listen(address)
}
