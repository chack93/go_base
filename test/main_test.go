package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/chack93/go_base/internal/domain"
	"github.com/chack93/go_base/internal/service/config"
	"github.com/chack93/go_base/internal/service/database"
	"github.com/chack93/go_base/internal/service/logger"
	"github.com/chack93/go_base/internal/service/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	if err := config.Init(); err != nil {
		logrus.Fatalf("config init failed, err: %v", err)
	}
	if err := logger.Init(); err != nil {
		logrus.Fatalf("log init failed, err: %v", err)
	}
	if err := database.New().Init(); err != nil {
		logrus.Fatalf("database init failed, err: %v", err)
	}
	if err := domain.Init(); err != nil {
		logrus.Fatalf("domain init failed, err: %v", err)
	}

	os.Exit(m.Run())
}

func Request(
	method string,
	path string,
	paramValues []string,
	body interface{},
) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	bodyJson, _ := json.Marshal(body)
	req := httptest.NewRequest(method, "/", bytes.NewReader(bodyJson))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
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

func CleanModelTS(a *model.Model, b *model.Model) {
	now := time.Now()
	a.CreatedAt = now
	b.CreatedAt = now
	a.UpdatedAt = now
	b.UpdatedAt = now
	a.DeletedAt.Time = now
	b.DeletedAt.Time = now
	a.DeletedAt.Valid = false
	b.DeletedAt.Valid = false
}
