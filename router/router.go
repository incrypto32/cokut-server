package router

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Tempelate
type Template struct {
	templates *template.Template
}

// Echo renderer interface implementation
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func New() *echo.Echo {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()

	// Echo renderer
	e.Renderer = t

	// Static files
	e.Static("/static", "assets")
	e.Static("/files", "files")

	// Not found handler
	echo.NotFoundHandler = func(c echo.Context) error {
		// Render your 404 page
		return c.String(http.StatusNotFound, "404 Not found")
	}

	// Middlewares
	e.Use(middleware.Recover())

	return e
}
