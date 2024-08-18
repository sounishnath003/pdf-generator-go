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
			Quantity:    generateRandomNumber(92, 160),
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
			ProductName: "Churidar Piece",
			Quantity:    generateRandomNumber(20, 70),
			Price:       generateRandomNumber(1500, 2300),
		},
		{
			ID:          4,
			ProductName: "Fancy Sarees",
			Quantity:    generateRandomNumber(60, 130),
			Price:       generateRandomNumber(1800, 2600),
		},
		{
			ID:          5,
			ProductName: "Churidar",
			Quantity:    generateRandomNumber(110, 180),
			Price:       generateRandomNumber(1300, 2300),
		},
		{
			ID:          6,
			ProductName: "Kurti",
			Quantity:    generateRandomNumber(60, 110),
			Price:       generateRandomNumber(500, 1280),
		},
		{
			ID:          7,
			ProductName: "Bed Sheets",
			Quantity:    generateRandomNumber(5, 13),
			Price:       generateRandomNumber(1000, 2000),
		},
		{
			ID:          8,
			ProductName: "Sweaters",
			Quantity:    generateRandomNumber(5, 13),
			Price:       generateRandomNumber(400,980),
		},
		{
			ID:          9,
			ProductName: "Winter Garments",
			Quantity:    generateRandomNumber(13, 30),
			Price:       generateRandomNumber(820, 1600),
		},
		{
			ID:          10,
			ProductName: "Cosmetics & Miscellaneous",
			Quantity:    generateRandomNumber(13, 30),
			Price:       generateRandomNumber(820, 1600),
		},
		{
			ID:          11,
			ProductName: "Others",
			Quantity:    generateRandomNumber(60,100),
			Price:       generateRandomNumber(2000, 4000),
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
	totalAmount := calculateTotalAmount(products)
	percentageLess := 25
	grossDiscount := float64(totalAmount) * float64(percentageLess) / 100.0
	amountAfterDiscount := float64(totalAmount) - (grossDiscount)

	acknowledgementRules := []models.Acknowlegement{
		{
			Norm:    "Sole Ownership and No Encumbrances",
			Details: "The assets listed in this statement are solely owned by the undersigned and are free from any liens, charges, or claims by any other party.",
		},
		{
			Norm:    "Accurate Inventory Valuation",
			Details: "The quantity and valuation of raw materials, work-in-progress, and finished goods as stated above represent a true and accurate reflection of the actual stock levels and their corresponding values as per the company's records.",
		},
		{
			Norm:    "Fair Market Valuation",
			Details: "The prices used to value the assets are the lowest of the market price, contract price, or controlled price, whichever is applicable.",
		},
		{
			Norm:    "Adequate Insurance Coverage",
			Details: "All assets listed in this statement are fully insured for their market value, and the relevant insurance policies are in force and held by the undersigned.",
		},
	}

	return &models.ReportGenerationModel{
		Name:           "sales_report",
		Date:           time.Now().Format("January-2006"),
		ShopName:       "MANANI",
		Propietor:      "Mrs. Manisha Nath",
		CashCreditBank: "State Bank of India - Ghatal Branch",
		Heading: models.HeaderInfo{
			Heading1: "Serial No. or the Stock Statement",
			Heading2: "CASH CREDIT ACCOUNT With",
			Heading3: "STATEMENT OF ASSETS / STOCKS HYPOTHDCATED AS NO",
		},
		PercentageLess:      int8(percentageLess),
		Products:            products,
		TotalAmount:         totalAmount,
		AmountAfterDiscount: amountAfterDiscount,
		GrossDiscountAmount: grossDiscount,
		AcknowlegmentRules:  acknowledgementRules,
		AuthorName:          "Sounish Nath",
	}
}

func calculateTotalAmount(products []models.MananiProduct) int {
	totalAmt := 0
	for _, product := range products {
		totalAmt += (product.Price * product.Quantity)
	}
	return totalAmt
}
