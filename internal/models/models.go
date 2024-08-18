package models

type IndexPageModel struct {
	Name   string `json:"name"`
	BadAss string `json:"badAss"`
	Date   string `json:"date"`
}

type HeaderInfo struct {
	Heading1 string `json:"heading_1"`
	Heading2 string `json:"heading_2"`
	Heading3 string `json:"heading_3"`
}

type ReportGenerationModel struct {
	Date           string     `json:"date"`
	Name           string     `json:"name"`
	ShopName       string     `json:"shopName"`
	CashCreditBank string     `json:"cashCreditBank"`
	Heading        HeaderInfo `json:"headerInfo"`
	PercentageLess int8       `json:"percentageLess"`
	AuthorName     string     `json:"authorName"`
}
