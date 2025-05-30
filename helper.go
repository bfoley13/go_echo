package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetGhost(c echo.Context) error {
	// test again
	// and again
	// another tet
	return c.String(http.StatusOK, "oooOOOoooOOOooo\n")
}
