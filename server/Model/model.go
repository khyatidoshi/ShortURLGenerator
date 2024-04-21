package model

type GenerateShortURLReq struct {
	LongURL    string `json:"long_url"`
	ExpiryDate string `json:"expiry,omitempty"`
}

type UrlData struct {
	tableName  struct{} `pg:"tinyurldata,alias:tinyurldata"`
	ShortURL   string   `json:"short_url" pg:"short_url,pk,on_delete:RESTRICT"`
	LongURL    string   `json:"long_url" pg:"long_url,notnull"`
	ExpiryDate int64    `json:"expiry,omitempty" pg:"expiry"`
}

// Access event details
type AccessDetails struct {
	tableName  struct{} `pg:"eventdetails,alias:eventdetails"`
	ShortURL   string   `json:"short_url" pg:"short_url,notnull,on_update:CASCADE"`
	AccessedAt int64    `json:"accessed_at" pg:"accessed_at,notnull"`
}

type Stats struct {
	DayCount   string `json:"day_count"`
	WeekCount  string `json:"week_count"`
	TotalCount int64  `json:"total_count"`
}
