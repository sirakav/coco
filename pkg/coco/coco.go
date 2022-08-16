package coco

import (
	_ "embed"

	"github.com/IncSW/geoip2"
	"github.com/cornelk/hashmap"
	"github.com/sams96/rgeo"
)

type COCO struct {
	Countries             []*CountryData
	Cities                []*CityData
	iso2ToCountryHM       *hashmap.HashMap[string, *CountryData]
	iso3ToCountryHM       *hashmap.HashMap[string, *CountryData]
	isoNumericToCountryHM *hashmap.HashMap[string, *CountryData]

	GEOLiteDB *geoip2.CountryReader
	RGEO      *rgeo.Rgeo
}

func NewCOCO() *COCO {
	co := loadCountryData()
	return &COCO{
		Countries:             co,
		Cities:                loadCityData(),
		iso2ToCountryHM:       getISO2ToCountryHM(co),
		iso3ToCountryHM:       getISO3ToCountryHM(co),
		isoNumericToCountryHM: getISONumericToCountryHM(co),

		GEOLiteDB: loadGEOLiteDatabase(),
		RGEO:      loadRGEO(),
	}
}
