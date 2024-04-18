package model

type GenerateShortURLReq struct {
	LongURL    string `json:"long_url"`
	ExpiryDate int64  `json:"expiry,omitempty"`
}

type UrlData struct {
	tableName  struct{} `pg:"tinyurldata,alias:tinyurldata"`
	ShortURL   string   `json:"short_url" pg:"short_url,pk"`
	LongURL    string   `json:"long_url" pg:"long_url,notnull"`
	ExpiryDate int64    `json:"expiry,omitempty" pg:"expiry"`
}

// AccessCounts represents access counts for a short URL
type AccessCounts struct {
	Last24Hours int `json:"last_24_hours"`
	LastWeek    int `json:"last_week"`
	AllTime     int `json:"all_time"`
}
