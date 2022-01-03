package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// NOTE : Using Fiber Web Framework

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		// NOTE : Inline handler function
		return c.SendString("Hello World!")
	})

	app.Get("/hello/:first_name/:last_name?", func(c *fiber.Ctx) error {
		// NOTE : Get params with optional + default
		fName := c.Params("first_name", "default_first_name")
		lName := c.Params("last_name", "default_last_name")
		return c.SendString(fmt.Sprintf("fn = %v | ln = %v", fName, lName))
	})

	app.Listen(":8000")

	_ = nativeMux()
}

func nativeMux() error {
	nativeApp := http.NewServeMux()

	nativeApp.HandleFunc("/native/mux/{id}", func(rw http.ResponseWriter, r *http.Request) {
		// NOTE : Manually Handle stuff....
		log.Println(r.Method)
	})

	// http.ListenAndServe(":8002", nativeApp)

	return nil
}
