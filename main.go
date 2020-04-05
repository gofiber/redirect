// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package redirect

import "github.com/gofiber/fiber"

// Config ...
type Config struct {
	Filter func(*fiber.Ctx) bool // Default: nil
	Rules  map[string]Rule       // The key we should be URL that you wish redirect
}

// Rule ...
type Rule struct {
	RedirectTo string
	StatusCode int // Default: 302
}

// New ...
func New(config ...Config) func(*fiber.Ctx) {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	statusCode := 302
	for _, v := range cfg.Rules {
		if v.StatusCode != statusCode {
			v.StatusCode = statusCode
		}
	}

	return func(c *fiber.Ctx) {
		if cfg.Filter != nil && cfg.Filter(c) {
			c.Next()
			return
		}
		for k, v := range cfg.Rules {
			if c.Path() == k {
				if v.RedirectTo != "" {
					c.Redirect(v.RedirectTo, v.StatusCode)
					return
				}
				location := c.Get("Location")
				if location != "" {
					c.Redirect(location, v.StatusCode)
					return
				}
			}
		}
		c.Next()
		return
	}
}
