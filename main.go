package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getGhost(c echo.Context) error {
	return c.String(http.StatusOK, "oooOOOoooOOOooo\n")
}

func getEcho(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("%s %s\n", c.Param("statement"), c.Param("statement")))
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/echo/:statement", getEcho)
	e.GET("*", getGhost)

	// adding comment for test
	// adding more test comments
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
