package fiberservertiming

import (
	"github.com/gofiber/fiber/v2"
	servertiming "github.com/mitchellh/go-server-timing"
)

type contextKeyType struct{}

var contextKey = contextKeyType(struct{}{})

func FromContext(c *fiber.Ctx) *servertiming.Header {
	return c.UserContext().Value(contextKey).(*servertiming.Header)
}

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Header is the header key where to get/set the unique request ID
	//
	// Optional. Default: "Server-Timing"
	Header string

	// AccessHeader is the header key to get/set access control origin for timing
	//
	// Optional. Default: "Timing-Allow-Origin"
	AccessHeader string

	// AllowOrigin access control origin for timing information
	//
	// Optional. Default: "*"
	AllowOrigins string
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:         nil,
	Header:       fiber.HeaderServerTiming,
	AccessHeader: fiber.HeaderTimingAllowOrigin,
	AllowOrigins: "*",
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Header == "" {
		cfg.Header = ConfigDefault.Header
	}
	if cfg.AccessHeader == "" {
		cfg.AccessHeader = ConfigDefault.AccessHeader
	}
	return cfg
}
