// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package redirect

import (
	"regexp"
	"strings"

	"github.com/gofiber/fiber"
)

// Config ...
type Config struct {
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter func(*fiber.Ctx) bool
	// Here you must set URL that you wishing to make the redirect with your rules.
	Rules map[string]Rule
}

// Rule ...
type Rule struct {
	// if you don't set it will to get of header location or refer
	RedirectTo string
	// Default: 302
	StatusCode  int
	hasWildcard bool
}

// New ...
func New(config ...Config) func(*fiber.Ctx) {
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	for k, v := range cfg.Rules {
		statusCode := 302
		if v.StatusCode > 0 {
			statusCode = v.StatusCode
		}
		cfg.Rules[k] = Rule{
			RedirectTo:  v.RedirectTo,
			StatusCode:  statusCode,
			hasWildcard: strings.Contains(k, "*"),
		}
	}

	return func(c *fiber.Ctx) {
		if cfg.Filter != nil && cfg.Filter(c) {
			c.Next()
			return
		}
		for k, v := range cfg.Rules {
			if c.Path() == k {
				redirectTo(v, c)
				return
			} else if v.hasWildcard {
				k = strings.Replace(k, "*", ".*", -1) + "$"
				re := regexp.MustCompile(k)
				if re.MatchString(c.Path()) {
					redirectTo(v, c)
					return
				}
			}
		}
		c.Next()
		return
	}
}

func redirectTo(rule Rule, c *fiber.Ctx) {
	location := c.Get("Location")
	if rule.RedirectTo != "" {
		c.Redirect(rule.RedirectTo, rule.StatusCode)
		return
	} else if location != "" {
		c.Redirect(location, rule.StatusCode)
		return
	}
}
