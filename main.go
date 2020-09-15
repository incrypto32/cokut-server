package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/routes"
	"github.com/incrypt0/cokut-server/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tempelate
type Template struct {
	templates *template.Template
}

// Echo renderer interface implementation
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	var PORT string
	args := os.Args[1:]

	fmt.Println("Welcome Hurray...!!!")

	// template init
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	// echo instance
	e := echo.New()

	// Connect to mongo
	services.ConnectMongo()

	// echo renderer
	e.Renderer = t

	// static files
	e.Static("/static", "assets")

	// middlewares
	e.Use(middleware.Recover())

	// not found handler
	echo.NotFoundHandler = func(c echo.Context) error {
		// render your 404 page
		return c.String(http.StatusNotFound, "404 Not found")
	}

	if _, err := services.InitFire(); err != nil {
		log.Panic(err)
	}

	// routes Handler
	routesHandler(e)

	// server

	if len(args) == 0 {

		PORT = os.Getenv("PORT")

		if PORT == "" {
			fmt.Println("PORT is empty")
			PORT = "4000"
		}

	} else {
		PORT = "4000"

	}

	e.Logger.Fatal(e.Start(":" + PORT))
}

func routesHandler(e *echo.Echo) {
	// routes
	e.GET("/", index)

	e.POST("/checkphone", checkUser)

	// Route groups
	api := e.Group("/api")
	admin := e.Group("/admin")

	// Route handlers
	routes.Api(api)
	routes.Admin(admin)

	// Basic Testing whether api works

}

// Index handler
func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "Hello,World")
}

// Get Users Details
func GetUserDetails(c echo.Context) (err error) {
	u := new(models.User)

	if err = c.Bind(u); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     err,
		})
	}

	u.ID = primitive.NewObjectID()

	id, err := models.InsertUser(u)

	if err != nil {
		return c.JSON(http.StatusExpectationFailed, echo.Map{
			"success": false,
			"msg":     err.Error(),
		})
	}

	return c.JSON(http.StatusExpectationFailed, echo.Map{
		"success": true,
		"id":      id,
	})
}

// Check if a user exists with phone and email
func checkUser(c echo.Context) (err error) {
	fmt.Println("Authentication : ", c.Request().Header.Get("Authentication"))

	phone := c.FormValue("phone")

	if models.CheckUser(phone) {
		return c.JSON(http.StatusAccepted, echo.Map{
			"success": true,
			"exist":   true,
		})
	}
	return c.JSON(http.StatusAccepted, echo.Map{
		"success": true,
		"exist":   false,
	})
}
