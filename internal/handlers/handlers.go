package handlers

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sounishnath003/pdf-generator-go/internal/htmx"
	"github.com/sounishnath003/pdf-generator-go/internal/models"
	"github.com/sounishnath003/pdf-generator-go/internal/service"
)

func HandleIndexPage(ctx echo.Context) error {
	data := models.IndexPageModel{
		Name:   "index",
		Date:   time.Now().Format("January-2006"),
		BadAss: "Batch-Sounish-Nath",
	}
	return ctx.Render(http.StatusOK, data.Name, &data)
}

var reportGenerationTempl = htmx.ParseAndRetrieveTemplFile("templates/layout.html.templ")

func HandleReportGeneration(ctx echo.Context) error {
	data := service.GenerateLatestSalesReport()

	ch := make(chan string)
	go generatePdf(ctx, data, ch)
	// Wait for the goroutine to finish using a channel
	data.DownloadPathUri = <-ch

	return ctx.Render(http.StatusOK, data.Name, &data)
}

// generatePdf generates a PDF file based on the provided data and sends the file path through the channel.
func generatePdf(ctx echo.Context, data *models.ReportGenerationModel, ch chan string) {
	// Create a new bytes buffer to hold the source HTML content.
	sourceBytes := bytes.NewBuffer([]byte{})

	// Execute the template with the provided data and write the result to the sourceBytes buffer.
	if err := reportGenerationTempl.ExecuteTemplate(sourceBytes, data.Name, data); err != nil {
		// Log the error if template execution fails.
		log.Println("an error occurred", err)
		// Send the error to the context.
		ctx.Error(err)
		return
	}

	// Generate a PDF file from the source HTML content.
	pdfFile, err := service.GeneratePdfFromSource(sourceBytes.Bytes())
	if err != nil {
		// Log the error if PDF generation fails.
		ctx.Logger().Info("error occurred", err)
		// Send the error to the context.
		ctx.Error(err)
		return
	}

	// Send the path of the generated PDF file through the channel.
	ch <- pdfFile
}
