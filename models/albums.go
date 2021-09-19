package models

// https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Albums []Album
