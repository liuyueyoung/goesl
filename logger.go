// Copyright 2015 Nevio Vesic
// Please check out LICENSE file for more information about what you CAN and what you CANNOT do!
// Basically in short this is a free software for you to do whatever you want to do BUT copyright must be included!
// I didn't write all of this code so you could say it's yours.
// MIT License

package goesl

import (
	"os"

	"github.com/op/go-logging"
)

var (
	log = logging.MustGetLogger("goesl")

	// Example format string. Everything except the message has a custom color
	// which is dependent on the log level. Many fields have a custom output
	// formatting too, eg. the time returns the hour down to the milli second.
	format = logging.MustStringFormatter(
		"%{color}%{time:15:04:05.000} %{shortfile} ▶ %{level:.8s}%{color:reset} %{message}",
	)
)

func SetLogLevel(l logging.Level) {
	logging.SetLevel(l, "goesl")
}

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(formatter)
}
