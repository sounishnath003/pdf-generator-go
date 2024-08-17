package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sounishnath003/pdf-generator-go/internal/models"
)

func HandleIndexPage(ctx echo.Context) error {
	data := models.IndexPageModel{
		Name:   "index",
		Date:   time.Now().Format("January-2006"),
		BadAss: "Batch-Sounish-Nath",
	}
	return ctx.Render(http.StatusOK, data.Name, &data)
}

func HandleReportGeneration(ctx echo.Context) error {
	data := models.ReportGenerationModel{
		Name:       "sales_report",
		AuthorName: "Sounish Nath",
	}
	return ctx.Render(http.StatusOK, data.Name, &data)
}
