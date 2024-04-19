package model

type GenerateShortURLReq struct {
	LongURL    string `json:"long_url"`
	ExpiryDate string `json:"expiry,omitempty"`
}

type UrlData struct {
	tableName  struct{} `pg:"tinyurldata,alias:tinyurldata"`
	ShortURL   string   `json:"short_url" pg:"short_url,pk"`
	LongURL    string   `json:"long_url" pg:"long_url,notnull"`
	ExpiryDate int64    `json:"expiry,omitempty" pg:"expiry"`
}

// Access event details
type AccessDetails struct {
	ShortURL        string `json:"short_url" pg:"short_url,notnull"`
	AccessTimestamp int64  `json:"expiry,omitempty" pg:"access_timestamp,notnull"`
	ResponseStatus  string `json:"response_status" pg:"response_status,notnull"`
}
