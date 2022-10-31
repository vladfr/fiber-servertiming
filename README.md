[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/vladfr/fiber-servertiming/v2) 

# Fiber Server Timing

This is a Fiber middleware for the [W3C Server-Timing API](see https://w3c.github.io/server-timing/) based on [mitchellh/go-server-timing](https://github.com/mitchellh/go-server-timing)

## Getting started

The Server-Timing API adds HTTP Headers for server-side performance metrics, such as time to render, DB calls latency, etc. You can view this data in the browser, and it can be collected by analytics providers via JavaScript.

This middleware adds a route timer by default, which should measure all the time spent in a request. You can use it to add your own custom timers via Fiber's user context.

## Usage

Import the middleware package that is part of the Fiber web framework

```
import (
  "github.com/gofiber/fiber/v2"
  "github.com/vladfr/fiber-servertiming/v2"
)
```

After you initiate your Fiber app, you can attach the middleware like this:

```
app := fiber.New()

# add the middleware; you should add it as the first middleware
# so you'll include the other middlewares in the final timing
app.Use(fiberservertiming.New())
```

To add metrics to this header, use Context:

```
app.Get("/server", func(c *fiber.Ctx) error {
  timing := fiberservertiming.FromContext(c)

  # time-consuming tasks

  # you should defer the .Stop() call so it triggers at the very end of this function
  defer timing.NewMetric("backendcall").WithDesc("My Description").Start().Stop()
}
```

## Metrics naming

In the Server-Timing spec, metric names cannot have any special characters, so use just letters and numbers. Also, since this is a http header, values for name and description should be kept as concise as possible.

## Security

Timing information is sent by the server in the form of a header. You should add additional protection if you don't want this data exposed in production, e.g. only enable this middleware on-demand or in dev or stage environments.

On the client-side, you can control which origins have access to the information via Javascript by setting the `Timing-Allow-Origin` header. This works exactly like a CORS header. See `/examples` for a full server/client example. This allows you to control which 3rd party scripts can access your timing data.

```
app := fiber.New()
app.Use(fiberservertiming.New(fiberservertiming.Config{
  AllowOrigins: "http://127.0.0.1:3000, http://127.0.0.1:3000/",
}))
```
