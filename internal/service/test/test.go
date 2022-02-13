package test

import (
	"io"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/chack93/go_base/internal/service/config"
	"github.com/chack93/go_base/internal/service/database"
	"github.com/chack93/go_base/internal/service/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Request(
	method string,
	path string,
	paramValues []string,
	body io.Reader,
) (echo.Context, *httptest.ResponseRecorder) {
	ensureDbConnection()

	e := echo.New()
	req := httptest.NewRequest(method, "/", body)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	var paramNames = []string{}
	for _, el := range strings.Split(path, "/") {
		if strings.Index(el, ":") == 0 {
			paramNames = append(paramNames, el[1:])
		}
	}
	ctx.SetPath(path)
	ctx.SetParamNames(paramNames...)
	ctx.SetParamValues(paramValues...)
	return ctx, rec
}

func CleanModel(a *model.Model, b *model.Model) {
	a.UUID = uuid.Nil
	b.UUID = uuid.Nil
	now := time.Now()
	a.CreatedAt = now
	b.CreatedAt = now
	a.UpdatedAt = now
	b.UpdatedAt = now
	a.DeletedAt.Time = now
	b.DeletedAt.Time = now
}

func ensureDbConnection() {
	if database.Get() != nil {
		return
	}
	if err := config.Init(); err != nil {
		logrus.Fatalf("config init failed, err: %v", err)
	}
	if err := database.New().Init(); err != nil {
		logrus.Fatalf("database init failed, err: %v", err)
	}
}
