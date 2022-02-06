package fiberservertiming

import servertiming "github.com/mitchellh/go-server-timing"

func newHeader() *servertiming.Header {
	var h servertiming.Header
	return &h
}
