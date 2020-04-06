// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber

package redirect

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber"
)

func Test_Redirect(t *testing.T) {
	app := *fiber.New()

	rules := map[string]Rule{
		"/default":     Rule{RedirectTo: "google.com", StatusCode: 301},
		"/default/*":   Rule{RedirectTo: "fiber.wiki", StatusCode: 307},
		"/pattern":     Rule{RedirectTo: "", StatusCode: 303},
		"/pattern/xyz": Rule{RedirectTo: "golang.org"},
	}

	app.Use(
		New(Config{Rules: rules}),
	)

	app.Get("/new", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	tests := []struct {
		name       string
		url        string
		redirectTo string
		statusCode int
	}{
		{
			name:       "should be returns status 302 without a wildcard",
			url:        "/default",
			redirectTo: "google.com",
			statusCode: 301,
		},
		{
			name:       "should be returns status 307 using wildcard",
			url:        "/default/xyz",
			redirectTo: "fiber.wiki",
			statusCode: 307,
		},
		{
			name:       "should be returns status 303 without set redirectTo to use the default",
			url:        "/pattern",
			redirectTo: "github.com/gofiber/redirect",
			statusCode: 303,
		},
		{
			name:       "should return the status code default",
			url:        "/pattern/xyz",
			redirectTo: "golang.org",
			statusCode: 302,
		},
		{
			name:       "access URL without rule",
			url:        "/new",
			statusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tt.url, nil)
			req.Header.Set("Location", "github.com/gofiber/redirect")
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf(`%s: %s`, t.Name(), err)
			}
			if resp.StatusCode != tt.statusCode {
				t.Fatalf(`%s: StatusCode: got %v - expected %v`, t.Name(), resp.StatusCode, tt.statusCode)
			}
			if resp.Header.Get("Location") != tt.redirectTo {
				t.Fatalf(`%s: Expecting Location: %s`, t.Name(), tt.redirectTo)
			}
		})
	}

}
