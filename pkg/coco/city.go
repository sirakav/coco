package coco

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strings"

	"database/sql"

	"github.com/gocarina/gocsv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirakav/coco/resources"
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
	GeonameID        int     `csv:"geonameid"`
	Name             string  `csv:"name"`
	ASCIIName        string  `csv:"asciiname"`
	AlternateNames   string  `csv:"alternatenames"`
	Latitude         float64 `csv:"latitude"`
	Longitude        float64 `csv:"longitude"`
	FeatureClass     string  `csv:"feature_class"`
	FeatureCode      string  `csv:"feature_code"`
	CountryCode      string  `csv:"country_code"`
	CC2              string  `csv:"cc2"`
	Admin1Code       string  `csv:"admin1_code"`
	Admin2Code       string  `csv:"admin2_code"`
	Admin3Code       string  `csv:"admin3_code"`
	Admin4Code       string  `csv:"admin4_code"`
	Population       int     `csv:"population"`
	Elevation        int     `csv:"elevation"`
	Dem              int     `csv:"dem"`
	Timezone         string  `csv:"timezone"`
	ModificationDate string  `csv:"modification_date"`
}

type CityDB struct {
	db *sql.DB
}

func NewCityDB() *CityDB {
	ciDB := &CityDB{}
	ciDB.db = ciDB.createMemDB()
	ciDB.createTables()
	ciDB.loadData()
	ciDB.createIndexes()

	return ciDB
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

	return c
}

func (c *CityDB) createIndexes() {
	_, err := c.db.Exec("CREATE INDEX name ON cities (name COLLATE NOCASE);")
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.db.Exec("CREATE INDEX asciiname ON cities (asciiname COLLATE NOCASE);")
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.db.Exec("CREATE INDEX alternatenames ON cities (alternatenames COLLATE NOCASE);")
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.db.Exec("CREATE INDEX population ON cities (population);")
	if err != nil {
		log.Fatal(err)
	}
}

func (c *CityDB) createTables() {
	sqlStmt := `
		CREATE TABLE cities(
			"geonameid" NUMERIC,
			"name" TEXT,
			"asciiname" TEXT,
			"alternatenames" TEXT,
			"latitude" NUMERIC,
			"longitude" NUMERIC,
			"feature_class" TEXT,
			"feature_code" TEXT,
			"country_code" TEXT,
			"cc2" TEXT,
			"admin1_code" TEXT,
			"admin2_code" TEXT,
			"admin3_code" TEXT,
			"admin4_code" TEXT,
			"population" INT,
			"elevation" INT,
			"dem" INT,
			"timezone" TEXT,
			"modification_date" TEXT
		);
	`

	_, err := c.db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *CityDB) loadData() {
	s := `INSERT INTO cities VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	for _, city := range loadCityData() {
		_, err := c.db.Exec(s, city.GeonameID, city.Name, city.ASCIIName, city.AlternateNames, city.Latitude, city.Longitude, city.FeatureClass, city.FeatureCode, city.CountryCode, city.CC2, city.Admin1Code, city.Admin2Code, city.Admin3Code, city.Admin4Code, city.Population, city.Elevation, city.Dem, city.Timezone, city.ModificationDate)
		if err != nil {
			log.Printf("%q: %s\n", err, s)
		}
	}
}

func (c *CityDB) createMemDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (c *CityDB) getCityDataFromQuery(q string, a ...any) []*CityData {
	var o []*CityData
	stmt, err := c.db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(a...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var city CityData
		if err := rows.Scan(&city.GeonameID, &city.Name, &city.ASCIIName, &city.AlternateNames, &city.Latitude,
			&city.Longitude, &city.FeatureClass, &city.FeatureCode, &city.CountryCode, &city.CC2, &city.Admin1Code,
			&city.Admin2Code, &city.Admin3Code, &city.Admin4Code, &city.Population, &city.Elevation, &city.Dem,
			&city.Timezone, &city.ModificationDate); err != nil {
			log.Fatal(err)
		}
		o = append(o, &city)
	}
	return o
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
	ci := c.CityDB.getCityDataFromQuery("SELECT * FROM cities WHERE name=? COLLATE NOCASE ORDER BY population DESC LIMIT 1;", name)
	if len(ci) > 0 {
		return ci[0], nil
	}
	if extendedSearch {
		ci = c.CityDB.getCityDataFromQuery("SELECT * FROM cities WHERE asciiname=? COLLATE NOCASE ORDER BY population DESC LIMIT 1;", name)
		if len(ci) > 0 {
			return ci[0], nil
		}
		ci = c.CityDB.getCityDataFromQuery("SELECT * FROM cities WHERE alternatenames LIKE ? COLLATE NOCASE ORDER BY population DESC LIMIT 1;", "%"+name+"%")
		if len(ci) > 0 {
			return ci[0], nil
		}
	}
	return &CityData{}, errors.New("city not found")
}
