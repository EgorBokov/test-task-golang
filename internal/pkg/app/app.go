package app

import (
	"fmt"
	"github.com/egorbokov/middleware/internal/app/endpoint"
	"github.com/egorbokov/middleware/internal/app/mw"
	"github.com/labstack/echo/v4"
)
import "github.com/egorbokov/middleware/internal/app/service"

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()

	a.echo.Use(mw.RoleCheckMiddleware)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (app *App) Run() error {
	fmt.Println("Server is listening!")

	if err := app.echo.Start(":8080"); err != nil {
		fmt.Println("ERROR!")
	}

	return nil
}
