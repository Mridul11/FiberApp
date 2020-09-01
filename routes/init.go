package routes

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber"
	"github.com/gofiber/websocket"
)

type Person struct {
	Name string `json:"name"`
	Age  int16  `json:"age"`
}

var p = []Person{ 
	Person{Name: "Mishra Mridul", Age: 26}, 
	Person{Name: "Vipul Mridul", Age: 21},
}

// Exposing all the app Endpoints
func Expose(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) {
		c.Render("index", fiber.Map{
			"Title":  "Hello, World!",
			"Person": p,
		})
	})

	app.Get("/about", func(c *fiber.Ctx) {
		c.Render("about", nil)
	})

	app.Get("/contact", func(c *fiber.Ctx) {
		c.Render("contact", nil)
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// Websocket logic
		for {
			mtype, msg, err := c.ReadMessage()
			if err != nil {
				break
			}
			log.Printf("Read: %s", msg)

			err = c.WriteMessage(mtype, msg)
			if err != nil {
				break
			}
		}
	}))

	app.Get("api/age", func(c *fiber.Ctx) {
		fmt.Println(p, c.Hostname())
		c.Status(302).JSON(fiber.Map{
			"Person": p,
		})
	})

	app.Get("api/age/:value", func(c *fiber.Ctx) {
		c.Send("Get Request value from the path variable: ", c.Params("value"))
	})
}
