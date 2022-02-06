# Server Timing

This is a Fiber middleware for the [W3C Server-Timing API](see https://w3c.github.io/server-timing/) based on [mitchellh/go-server-timing](https://github.com/mitchellh/go-server-timing)

## Getting started

The Server-Timing API adds HTTP Headers for server-side performance metrics, such as time to render, DB calls latency, etc. You can view this data in the browser, and it can be collected by analytics providers via the browser.

This middleware adds a general request timer by default. You should use it to add your own custom timers via Fiber's user context.

## Usage

Import the middleware package that is part of the Fiber web framework

```
import (
  "github.com/gofiber/fiber/v2"
  "github.com/vladfr/fiber-server-timing"
)
```

After you initiate your Fiber app, you can attach the middleware like this:

```
fiber.Use(servertiming)
```
