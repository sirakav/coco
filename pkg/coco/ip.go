package coco

import (
	"net"

	"github.com/sirakav/coco/resources"

	"github.com/IncSW/geoip2"
)

func loadGEOLiteDatabase() *geoip2.CountryReader {
	reader, err := geoip2.NewCountryReader(resources.GEOLite2CountryMMDB)
	if err != nil {
		panic(err)
	}

	return reader
}

func (c *COCO) IPToCountry(ip string) (*CountryData, error) {
	record, err := c.GEOLiteDB.Lookup(net.ParseIP(ip))
	if err != nil {
		return &CountryData{}, err
	}
	return c.matchISO2ToCountry(record.Country.ISOCode)
}

func (c *COCO) IPToISO2(ip string) (string, error) {
	record, err := c.GEOLiteDB.Lookup(net.ParseIP(ip))
	if err != nil {
		return "", err
	}
	return record.Country.ISOCode, nil
}

func (c *COCO) IPToISO3(ip string) (string, error) {
	record, err := c.GEOLiteDB.Lookup(net.ParseIP(ip))

	if err != nil {
		return "", err
	}

	co, err := c.matchISO2ToCountry(record.Country.ISOCode)
	if err != nil {
		return "", err
	}
	return co.ISO3, nil
}
