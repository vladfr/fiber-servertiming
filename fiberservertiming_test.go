package fiberservertiming

import (
	"fmt"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	servertiming "github.com/mitchellh/go-server-timing"
)

func Test_Simple(t *testing.T) {
	app := fiber.New()

	app.Use(New())

	app.Get("/timing", func(c *fiber.Ctx) error {
		time.Sleep(12 * time.Millisecond)
		return c.SendStatus(fiber.StatusOK)
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/timing", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, "route;desc=\"/timing\";dur=", resp.Header.Get(fiber.HeaderServerTiming)[:25])
}

func Test_Multiple(t *testing.T) {
	app := fiber.New()

	app.Use(New())

	app.Get("/timing", func(c *fiber.Ctx) error {
		time.Sleep(12 * time.Millisecond)
		return c.SendStatus(fiber.StatusOK)
	})

	r := regexp.MustCompile("^route;desc=\"/timing\";dur=[0-9.,]{5,12}$")
	for i := 1; i < 5; i++ {
		resp, err := app.Test(httptest.NewRequest("GET", "/timing", nil))
		utils.AssertEqual(t, nil, err)
		headerStr := resp.Header.Get(fiber.HeaderServerTiming)
		match := r.MatchString(headerStr)
		utils.AssertEqual(t, match, true, fmt.Sprintf("regex does not match '%s'", headerStr))
	}
}

func Test_CustomMetric(t *testing.T) {
	app := fiber.New()

	app.Use(New())

	app.Get("/metric", func(c *fiber.Ctx) error {
		timing := FromContext(c)
		time.Sleep(10 * time.Millisecond)

		defer timing.NewMetric("backendcall").Start().Stop()
		time.Sleep(10 * time.Millisecond)

		return c.SendStatus(fiber.StatusOK)
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/metric", nil))
	utils.AssertEqual(t, nil, err)

	headerStr := resp.Header.Get(fiber.HeaderServerTiming)
	h, errh := servertiming.ParseHeader(headerStr)
	utils.AssertEqual(t, nil, errh)

	utils.AssertEqual(t, "route", h.Metrics[0].Name)
	utils.AssertEqual(t, "backendcall", h.Metrics[1].Name)
}

func Test_CustomMetricRegex(t *testing.T) {
	app := fiber.New()

	app.Use(New())

	app.Get("/metric", func(c *fiber.Ctx) error {
		timing := FromContext(c)
		time.Sleep(10 * time.Millisecond)

		defer timing.NewMetric("backendcall").Start().Stop()
		time.Sleep(10 * time.Millisecond)

		return c.SendStatus(fiber.StatusOK)
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/metric", nil))
	utils.AssertEqual(t, nil, err)

	headerStr := resp.Header.Get(fiber.HeaderServerTiming)
	match, errRe := regexp.MatchString("route;desc=\"/metric\";dur=[0-9.,]+,backendcall;dur=[0-9.,]", headerStr)
	utils.AssertEqual(t, nil, errRe)
	utils.AssertEqual(t, match, true, headerStr)
}
