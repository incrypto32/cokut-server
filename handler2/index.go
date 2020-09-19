package handler2

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) routesHandler(e *echo.Echo) {

	// routes
	e.GET("/", h.index)

}

// Index handler
func (h *Handler) index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "Hello,World")
}
