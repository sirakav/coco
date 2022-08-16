package coco

import (
	"log"

	"github.com/sams96/rgeo"
)

func loadRGEO() *rgeo.Rgeo {
	// Can be changed to rgeo.Countries10 for perfomance improvements
	r, err := rgeo.New(rgeo.Countries110)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func (c *COCO) LatLonToCountry(lat, lon float64) (*CountryData, error) {
	l, err := c.RGEO.ReverseGeocode([]float64{lat, lon})
	if err != nil {
		return &CountryData{}, err
	}
	return c.matchISO2ToCountry(l.CountryCode2)
}

func (c *COCO) LatLonToISO2(lat, lon float64) (string, error) {
	l, err := c.RGEO.ReverseGeocode([]float64{lat, lon})
	if err != nil {
		return "", err
	}
	return l.CountryCode2, nil
}

func (c *COCO) LatLonToISO3(lat, lon float64) (string, error) {
	l, err := c.RGEO.ReverseGeocode([]float64{lat, lon})
	if err != nil {
		return "", err
	}
	return l.CountryCode3, nil
}
