package sessions

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

// sessions middleware for Echo
func Sessions(name string, store sessions.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(contextKey, &session{ctx: c, name: name, store: store})
			return next(c)
		}
	}
}
