package session

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandlerCreate(c echo.Context) error {
	session := Session{}
	return c.JSON(http.StatusOK, session)
}

func HandlerList(c echo.Context) error {
	sessionList := []Session{}
	return c.JSON(http.StatusOK, sessionList)
}

func HandlerRead(c echo.Context) error {
	// id := c.Param("id")
	var session Session
	return c.JSON(http.StatusOK, session)
}

func HandlerUpdate(c echo.Context) error {
	// id := c.Param("id")
	var session Session
	return c.JSON(http.StatusOK, session)
}

func HandlerDelete(c echo.Context) error {
	// id := c.Param("id")
	var session Session
	return c.JSON(http.StatusOK, session)
}
