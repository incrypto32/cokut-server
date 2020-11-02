package middleware

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// FireAuthMiddleware is a middlerware which handles the authorisation checking for users.
func AdminCheckMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			adminUID := os.Getenv("COKUT_ADMIN_UID")

			if c.Get("uid") == adminUID {
				c.Set("admin", true)
			}

			if c.Get("admin").(bool) {
				return next(c)
			}

			return c.JSON(http.StatusExpectationFailed, echo.Map{"success": false, "msg": "Not Authorized"})
		}
	}
}
