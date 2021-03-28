package cli

type GeoRepo interface {
	Locate() (*Location, error)
}

type LocationLabel string

type Location struct {
	Country     string  `json:"country_name"`
	CountryCode string  `json:"country_code"`
	Region      string  `json:"region_name"`
	ZipCode     string  `json:"zip_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	City        string  `json:"city"`
}
