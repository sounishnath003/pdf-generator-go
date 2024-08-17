package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sounishnath003/pdf-generator-go/internal/htmx"
)

const (
	PublicHost = "localhost"
	Port       = "3000"
)

func main() {
	e := echo.New()

	// Little bit of middleware housekeeping
	e.Use(middleware.CSRF())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	// Initializes the template language
	htmx.NewTemplateRender(e, "public/*.html")

	// Create single SSR endpoints to generate the Pdf Reports
	e.GET("/", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "index", map[string]string{
			"BadAss": "Batch",
		})
	})

	e.GET("/generate-report", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "sales_report", map[string]string{
			"AuthorName": "Sounish Nath",
		})
	})

	log.Printf("server is up and running on port http://%s:%s/\n", PublicHost, Port)
	e.Logger.Panic(e.Start(fmt.Sprintf(":%s", Port)))
}
