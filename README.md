### Install
```
go get -u github.com/gofiber/fiber
go get -u github.com/gofiber/redirect
```
### Example
```go
package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/redirect"
)

func main() {
	app := fiber.New()
  
  // Optional
	config := redirect.Config{
		Handler: func(c *fiber.Ctx, err error) {
			c.SendString(err.Error())
			c.SendStatus(500)
		},
	}

	app.Use(redirect.New(config))

	app.Get("/newpath", func(c *fiber.Ctx) {
		c.Send("Hello!")
	})

	app.Listen(3000)
}
```
### Test
```curl
curl http://localhost:3000/oldpath
```
