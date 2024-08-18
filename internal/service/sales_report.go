package service

import (
	"time"

	"github.com/sounishnath003/pdf-generator-go/internal/models"
)

func GenerateLatestSalesReport() *models.ReportGenerationModel {
	return &models.ReportGenerationModel{
		Name:           "sales_report",
		Date:           time.Now().Format("January-2006"),
		ShopName:       `"MANANI"`,
		CashCreditBank: "State Bank of India - Ghatal Branch",
		Heading: models.HeaderInfo{
			Heading1: "Serial No. or the Stock Statement",
			Heading2: "CASH CREDIT ACCOUNT With",
			Heading3: "STATEMENT OF ASSETS / STOCKS HYPOTHDCATED AS NO",
		},
		PercentageLess: 25,
		AuthorName:     "Sounish Nath",
	}
}
