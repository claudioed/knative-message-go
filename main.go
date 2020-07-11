package main

import (
"net/http"
	"os"

	"github.com/labstack/echo"
"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/message", func(c echo.Context) error {
		message := os.Getenv("MESSAGE")
		if len(message) == 0{
			message = "v1"
		}
		return c.String(http.StatusOK,message)
	})

	e.GET("/health/live", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/health/ready", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
