package sessions

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

// sessions middleware for Echo
func Sessions(name string, store sessions.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			s := &session{ctx: c, name: name, store: store}
			c.Set(ContextKey, s)
			return next(c)
		}
	}
}
