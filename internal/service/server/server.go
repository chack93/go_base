package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/chack93/go_base/internal/domain/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Server struct {
	echo *echo.Echo
}

var server *Server

func Get() *Server {
	return server
}

func New() *Server {
	server = &Server{}
	return server
}

func (srv *Server) Init(wg *sync.WaitGroup) error {
	srv.echo = echo.New()
	srv.echo.HideBanner = true
	srv.echo.HidePort = true
	srv.echo.Use(middleware.Logger())
	srv.echo.Use(middleware.Recover())

	apiAppGroup := srv.echo.Group("/api/go_base")
	apiAppGroup.POST("/session/", session.HandlerCreate)
	apiAppGroup.GET("/session/", session.HandlerList)
	apiAppGroup.GET("/session/:id", session.HandlerRead)
	apiAppGroup.PUT("/session", session.HandlerUpdate)
	apiAppGroup.DELETE("/session/:id", session.HandlerDelete)

	address := fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))
	go func() {
		if err := srv.echo.Start(address); err != nil && err != http.ErrServerClosed {
			logrus.Warnf("server start failed, err: %v", err)
			wg.Done()
		}
	}()

	defer func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.echo.Shutdown(ctx); err != nil {
			logrus.Errorf("server shutdown failed, err: %v", err)
		}
		logrus.Info("server shutdown")
		wg.Done()
	}()

	for _, el := range srv.echo.Routes() {
		lastSlash := strings.LastIndex(el.Name, "/")
		domainHandler := el.Name[lastSlash:]
		logrus.Infof("%6s %s -> %s", el.Method, el.Path, domainHandler)
	}
	logrus.Infof("http server started on %s", address)
	return nil
}
