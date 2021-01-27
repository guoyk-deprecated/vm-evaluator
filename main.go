package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

var (
	optAddr string
)

type zeroReader struct {
}

func (z zeroReader) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0
	}
	n = len(p)
	return
}

func main() {
	flag.StringVar(&optAddr, "addr", ":80", "listen address")
	flag.Parse()

	log.Println("listening at", optAddr)

	e := echo.New()
	e.HidePort = true
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/zero", func(c echo.Context) error {
		return c.Stream(http.StatusOK, "application/octet-stream", zeroReader{})
	})
	e.Logger.Fatal(e.Start(optAddr))
}
