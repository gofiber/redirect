// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package redirect

import (
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber"
)

// Config ...
type Config struct {
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter func(*fiber.Ctx) bool
	// Handler is called when a panic occurs
	// Optional. Default: c.SendStatus(500)
	Handler func(*fiber.Ctx, error)
	// Log all errors to output
	// Optional. Default: false
	Log bool
	// Output is a writter where logs are written
	// Default: os.Stderr
	Output io.Writer
}

// New ...
func New(config ...Config) func(*fiber.Ctx) {
	// Init config
	var cfg Config
	// Set config if provided
	if len(config) > 0 {
		cfg = config[0]
	}
	// Set config default values
	if cfg.Handler == nil {
		cfg.Handler = func(c *fiber.Ctx, err error) {
			c.SendStatus(500)
		}
	}
	if cfg.Output == nil {
		cfg.Output = os.Stderr
	}
	// Return middleware handle
	return func(c *fiber.Ctx) {
		// Filter request to skip middleware
		if cfg.Filter != nil && cfg.Filter(c) {
			c.Next()
			return
		}
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				if cfg.Log {
					cfg.Output.Write([]byte(err.Error() + "\n"))
				}
				cfg.Handler(c, err)
			}
		}()
		c.Next()
	}
}
