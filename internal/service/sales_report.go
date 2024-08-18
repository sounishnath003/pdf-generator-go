package service

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/sounishnath003/pdf-generator-go/internal/models"
)

// generateRandomNumber generates a random number between start and end range.
func generateRandomNumber(start, end int) int {
	seedNo, err := strconv.Atoi(time.Now().Format("200601"))
	if err != nil {
		log.Println("an error occurred", err)
		return end
	}
	rand.NewSource(int64(seedNo))

	return rand.Intn(end-start) + start
}

// Generate Product catalog as per the current month sales.
func getProductCatalog() ([]models.MananiProduct, error) {
	return []models.MananiProduct{
		{
			ID:          1,
			ProductName: "Saree",
			Quantity:    generateRandomNumber(20, 70),
			Price:       generateRandomNumber(810, 1900),
		},
		{
			ID:          2,
			ProductName: "Saree",
			Quantity:    generateRandomNumber(20, 70),
			Price:       generateRandomNumber(1500, 2300),
		},
		{
			ID:          3,
			ProductName: "Churidar pics",
			Quantity:    generateRandomNumber(20, 70),
			Price:       generateRandomNumber(1500, 2300),
		},
		{
			ID:          4,
			ProductName: "Womens' Jens",
			Quantity:    generateRandomNumber(20, 70),
			Price:       generateRandomNumber(1500, 2300),
		},
	}, nil
}

// GenerateLatestSalesReport considering the Report and creating the data models.
//
// Returns the required information for the Templ template to showcase.
func GenerateLatestSalesReport() *models.ReportGenerationModel {
	products, err := getProductCatalog()
	if err != nil {
		panic(err)
	}
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
		Products:       products,
		TotalAmount:    calculateTotalAmount(products),
		AuthorName:     "Sounish Nath",
	}
}

func calculateTotalAmount(products []models.MananiProduct) int {
	totalAmt := 0
	for _, product := range products {
		totalAmt += (product.Price * product.Quantity)
	}
	return totalAmt
}
