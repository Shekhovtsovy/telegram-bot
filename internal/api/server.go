package api

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func NewServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	return e
}
