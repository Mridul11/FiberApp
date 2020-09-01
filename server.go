package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/template/html"
	router "./routes"
)

func main() {
	var app = fiber.New()

	engine := html.New("./views", ".html")
	// engine.Reload(true)
	// After you created your engine, you can pass it to Fiber's Views Engine
	app = fiber.New(&fiber.Settings{
		Views: engine,
	})
	app.Use(middleware.Logger())
	router.Expose(app)
	app.Listen(8000)
}
