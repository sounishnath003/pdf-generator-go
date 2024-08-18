package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sounishnath003/pdf-generator-go/internal/handlers"
	"github.com/sounishnath003/pdf-generator-go/internal/htmx"
)

var Envs = InitEnvs()

func main() {
	Init()

	log.Printf("loading.env:%v\n", Envs)
	e := echo.New()

	// Serve static assets
	e.Static("export", "temp")
	e.Static("public", "public")

	// Little bit of middleware housekeeping
	e.Use(middleware.CSRF())
	e.Pre(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	// Initializes the template language
	htmx.NewTemplateRender(e, "public/*.html")

	// Create single SSR endpoints to generate the Pdf Reports
	e.GET("/", handlers.HandleIndexPage)
	// Report Generation Endpoint
	e.GET("/generate-report", handlers.HandleReportGeneration)

	log.Printf("server is up and running on port http://%s:%s/\n", Envs.PublicHost, Envs.Port)
	e.Logger.Panic(e.Start(fmt.Sprintf(":%s", Envs.Port)))
}

// Init initializes the necessary directory structure for public assets exports
func Init() {
	err := os.MkdirAll("temp", 0o755)
	if err != nil {
		panic(err)
	}
}
