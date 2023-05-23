package mw

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func RoleCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		str := ctx.Request().Header.Get("User-Role")

		if str == "admin" {
			fmt.Println("red button found!")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
