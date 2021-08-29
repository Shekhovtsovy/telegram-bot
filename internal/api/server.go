package api

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

// NewServer returns new web server (using at least for prometheus metrics)
func NewServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	return e
}
