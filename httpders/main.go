package main

import(
	"github.com/gofiber/fiber/v2"

)

// net/http  (gin,echo,gorilla/mux)
// calyala/fasthttp

func main()  {
	
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	app.Get("/:key/*",ProxyHandler)

    app.Listen(":3000")

}