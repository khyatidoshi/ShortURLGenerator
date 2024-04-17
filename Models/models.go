package models

// ShortURL represents a short URL entry
type ShortURL struct {
	Short  string `json:"short"`
	Long   string `json:"long"`
	Expiry int64  `json:"expiry,omitempty"`
	// Add more fields as needed (creation time, etc.)
}

// AccessCounts represents access counts for a short URL
type AccessCounts struct {
	Last24Hours int `json:"last_24_hours"`
	LastWeek    int `json:"last_week"`
	AllTime     int `json:"all_time"`
}
