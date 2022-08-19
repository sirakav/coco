package coco

import (
	"github.com/sirakav/coco/resources"
	"encoding/csv"
	"io"
	"sort"

	"strings"

	"github.com/gocarina/gocsv"
)

// List of strings
type StringList struct {
	List []string
}

func (sl *StringList) UnmarshalCSV(csv string) (err error) {
	sl.List = strings.Split(csv, ",")
	return nil
}

func (sl *StringList) MarshalCSV() (string, error) {
	return strings.Join(sl.List, ","), nil
}

// The main 'geoname' table has the following fields :
// ---------------------------------------------------
// geonameid         : integer id of record in geonames database
// name              : name of geographical point (utf8) varchar(200)
// asciiname         : name of geographical point in plain ascii characters, varchar(200)
// alternatenames    : alternatenames, comma separated, ascii names automatically transliterated, convenience attribute from alternatename table, varchar(10000)
// latitude          : latitude in decimal degrees (wgs84)
// longitude         : longitude in decimal degrees (wgs84)
// feature class     : see http://www.geonames.org/export/codes.html, char(1)
// feature code      : see http://www.geonames.org/export/codes.html, varchar(10)
// country code      : ISO-3166 2-letter country code, 2 characters
// cc2               : alternate country codes, comma separated, ISO-3166 2-letter country code, 200 characters
// admin1 code       : fipscode (subject to change to iso code), see exceptions below, see file admin1Codes.txt for display names of this code; varchar(20)
// admin2 code       : code for the second administrative division, a county in the US, see file admin2Codes.txt; varchar(80)
// admin3 code       : code for third level administrative division, varchar(20)
// admin4 code       : code for fourth level administrative division, varchar(20)
// population        : bigint (8 byte int)
// elevation         : in meters, integer
// dem               : digital elevation model, srtm3 or gtopo30, average elevation of 3”x3” (ca 90mx90m) or 30”x30” (ca 900mx900m) area in meters, integer. srtm processed by cgiar/ciat.
// timezone          : the iana timezone id (see file timeZone.txt) varchar(40)
// modification date : date of last modification in yyyy-MM-dd format
type CityData struct {
	GeonameID        int        `csv:"geonameid"`
	Name             string     `csv:"name"`
	ASCIIName        string     `csv:"asciiname"`
	AlternateNames   StringList `csv:"alternatenames"`
	Latitude         float64    `csv:"latitude"`
	Longitude        float64    `csv:"longitude"`
	FeatureClass     string     `csv:"feature_class"`
	FeatureCode      string     `csv:"feature_code"`
	CountryCode      string     `csv:"country_code"`
	CC2              string     `csv:"cc2"`
	Admin1Code       string     `csv:"admin1_code"`
	Admin2Code       string     `csv:"admin2_code"`
	Admin3Code       string     `csv:"admin3_code"`
	Admin4Code       string     `csv:"admin4_code"`
	Population       int        `csv:"population"`
	Elevation        int        `csv:"elevation"`
	Dem              int        `csv:"dem"`
	Timezone         string     `csv:"timezone"`
	ModificationDate string     `csv:"modification_date"`
}

func loadCityData() []*CityData {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = '\t'
		return r
	})

	c := []*CityData{}
	if err := gocsv.UnmarshalBytes(resources.CityDataTSV, &c); err != nil {
		panic(err)
	}

	// Sort by population because we only return the cities with the highest population
	sort.Slice(c, func(i, j int) bool {
		return c[i].Population > c[j].Population
	})
	return c
}

// Returns city with highest population for the given city name
// Perfomance scales non linearly as first we try to match the city name
// than the ASCIIName and finally the AlternateNames which contains
// city names in other languages, this allows us to have multiple language support
// for the price of performance.
// ---------------------------------------------------
// Arguments:
// name - the name of the city
// extendedSearch - if true, we will search for the city in the alternate names
func (c *COCO) CityNameToCity(name string, extendedSearch bool) (*CityData, error) {
	for _, ci := range c.Cities {
		if strings.EqualFold(ci.Name, name) {
			return ci, nil
		}
	}

	if !extendedSearch {
		return &CityData{}, &CityNotFoundError{}
	}

	for _, ci := range c.Cities {
		if strings.EqualFold(ci.ASCIIName, name) {
			return ci, nil
		}
	}
	for _, ci := range c.Cities {
		for _, an := range ci.AlternateNames.List {
			if strings.EqualFold(an, name) {
				return ci, nil
			}
		}
	}
	return &CityData{}, &CityNotFoundError{}
}
