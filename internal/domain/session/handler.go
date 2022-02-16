package session

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func HandlerCreate(c echo.Context) error {
	request := Session{
		JoinCode: uuid.NewString(),
	}
	if err := CreateSession(&request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete")
	}
	return c.JSON(http.StatusOK, request)
}

func HandlerList(c echo.Context) error {
	var responseList = []Session{}
	if err := ListSession(&responseList); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete")
	}
	return c.JSON(http.StatusOK, responseList)
}

func HandlerRead(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad id, expected format: uuid")
	}
	var response Session
	if err := ReadSession(uuid, &response); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete")
	}
	return c.JSON(http.StatusOK, response)
}

func HandlerUpdate(c echo.Context) error {
	var request Session
	if err := c.Bind(&request); err != nil {
		logrus.Infof("bind body failed: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad body, expected format: Session.json")
	}
	err := UpdateSession(&request)
	if err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update")
	}
	return c.JSON(http.StatusOK, request)
}

func HandlerDelete(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad id, expected format: uuid")
	}
	var response Session
	err = DeleteSession(id, &response)
	if err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete")
	}
	return c.JSON(http.StatusOK, response)
}
