package middlewares

import (
	"github.com/lovego/goa/middlewares"
)

var Logger = middlewares.NewLogger(config.HttpLogger())
