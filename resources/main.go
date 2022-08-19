package resources

import (
	_ "embed"
)

var (
	//go:embed country_data.csv
	CountryDataCSV []byte

	//go:embed geolite2/GeoLite2-Country.mmdb
	GEOLite2CountryMMDB []byte

	//go:embed cities500.tsv
	CityDataTSV []byte
)
