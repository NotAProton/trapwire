package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

var loc, _ = time.LoadLocation("Asia/Kolkata")

func checkAuth(c echo.Context) bool {
	return c.Request().Header.Get("Authorization") == os.Getenv("AUTH")
}

func redir(c echo.Context) error {
	t := time.Now().In(loc).Format("15:04:05")
	dest, found := getDest(c.Param("wire"))
	if !found {
		insertLog(t, c.RealIP(), c.Request().UserAgent(), c.Request().Referer(), c.Param("wire"), "Not found")

		return c.Redirect(http.StatusTemporaryRedirect, "https://www.google.com")
	}
	insertLog(t, c.RealIP(), c.Request().UserAgent(), c.Request().Referer(), c.Param("wire"), dest)

	return c.Redirect(http.StatusTemporaryRedirect, dest)
}

func addWire(c echo.Context) error {
	if !checkAuth(c) {
		return c.NoContent(http.StatusUnauthorized)
	}
	wire := c.FormValue("wire")
	dest := c.FormValue("dest")
	if wire == "" || dest == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	insertWire(wire, dest)
	return c.NoContent(http.StatusOK)
}

func getLogs(c echo.Context) error {
	if !checkAuth(c) {
		return c.NoContent(http.StatusUnauthorized)
	}
	return c.JSON(http.StatusOK, getLogsData())
}
