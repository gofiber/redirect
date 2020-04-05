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
  
  app.Use(redirect.New(redirect.Config{
    Rules: map[string]string{
      "/old":   redirect.Rule{RedirectTo: "google.com", StatusCode: 301},
      "/old/*": redirect.Rule{RedirectTo: "fiber.wiki", StatusCode: 307},
    },
  }))
  
  app.Get("/new", func(c *fiber.Ctx) {
    c.Send("Hello, World!")
  })
  app.Get("/new/*", func(c *fiber.Ctx) {
    c.Send("Wildcard: ", c.Params("*"))
  })
  
  app.Listen(3000)
}

```

### Test
```curl
curl http://localhost:3000/old
curl http://localhost:3000/old/hello
```