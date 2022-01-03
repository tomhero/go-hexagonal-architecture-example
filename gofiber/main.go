package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// NOTE : Using Fiber Web Framework

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/hello/string", func(c *fiber.Ctx) error {
		// NOTE : Inline handler function
		return c.SendString("Hello World!")
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
