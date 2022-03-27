package fiberservertiming

import (
	"context"

	"github.com/gofiber/fiber/v2"
	servertiming "github.com/mitchellh/go-server-timing"
)

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config...)

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		timing := newHeader()
		c.SetUserContext(context.WithValue(c.UserContext(), contextKey, timing))

		defer addHeaders(cfg, c, timing)
		defer timing.NewMetric("latency").Start().Stop()
		return c.Next()
	}
}

func addHeaders(cfg Config, c *fiber.Ctx, h *servertiming.Header) {
	c.Append(cfg.AccessHeader, "Test")
	c.Append(cfg.Header, h.String())
}
