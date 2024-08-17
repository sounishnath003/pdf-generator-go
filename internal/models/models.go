package models

type IndexPageModel struct {
	Name   string `json:"name"`
	BadAss string `json:"badAss"`
	Date   string `json:"date"`
}

type ReportGenerationModel struct {
	Name       string `json:"name"`
	AuthorName string `json:"authorName"`
}
