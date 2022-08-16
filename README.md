# COCO
[![Codecov](https://img.shields.io/codecov/c/github/sirakav/coco?logo=codecov&style=for-the-badge)](https://codecov.io/gh/sirakav/coco)
[![](https://goreportcard.com/badge/github.com/sirakav/coco?style=for-the-badge)](https://goreportcard.com/report/github.com/sirakav/coco)
[![Release](https://img.shields.io/github/tag/sirakav/coco.svg?label=release&color=24B898&logo=github&style=for-the-badge)](https://github.com/sirakav/coco/releases/latest)

COCO, a GoLang library, strives to be fast and simple solution, for converting country or city names and codes, IP addresses and coordinates to standartized objects that represent and deliver usable country or city information.

## Key Features
Get country details from:
- Variuos country name formats or codes
- IP addresses
- Coordinates

Get city details from:
- Names

Every step is done offline to increase performance and decrease reliance on variuos online services.
This enables you to parse, validate or format huge amounts of data offline. 

## Usage
### Basic usage
```go
var cc = NewCOCO()
co, err := cc.CountryNameToCountry("Dem. People's Rep. of Korea", "")
if err != nil {
    panic(err)
}
fmt.PrintLn(co)
```

### Coordinates to country 
```go
var cc = NewCOCO()
s, err := cc.LatLonToISO2(60.1551229, 37.2166344)
if err != nil {
    panic(err)
}
fmt.PrintLn(s)
```

Further documentation TBD!

## Data structures
### Country data
```go
type CountryData struct {
	NameShort    string        `csv:"name_short"`
	NameOfficial string        `csv:"name_official"`
	Regex        CompiledRegex `csv:"regex"`
	ISO2         string        `csv:"ISO2"`
	ISO3         string        `csv:"ISO3"`
	ISOnumeric   string        `csv:"ISOnumeric"`
	UNcode       string        `csv:"UNcode"`
	FAOcode      string        `csv:"FAOcode"`
	GBDcode      string        `csv:"GBDcode"`
	Continent    string        `csv:"continent"`
	UNregion     string        `csv:"UNregion"`
	EXIO1        string        `csv:"EXIO1"`
	EXIO2        string        `csv:"EXIO2"`
	EXIO3        string        `csv:"EXIO3"`
	EXIO1_3L     string        `csv:"EXIO1_3L"`
	EXIO2_3L     string        `csv:"EXIO2_3L"`
	EXIO3_3L     string        `csv:"EXIO3_3L"`
	WIOD         string        `csv:"WIOD"`
	Eora         string        `csv:"Eora"`
	MESSAGE      string        `csv:"MESSAGE"`
	IMAGE        string        `csv:"IMAGE"`
	REMIND       string        `csv:"REMIND"`
	OECD         string        `csv:"OECD"`
	EU           string        `csv:"EU"`
	EU28         string        `csv:"EU28"`
	EU27         string        `csv:"EU27"`
	EU27_2007    string        `csv:"EU27_2007"`
	EU25         string        `csv:"EU25"`
	EU15         string        `csv:"EU15"`
	EU12         string        `csv:"EU12"`
	EEA          string        `csv:"EEA"`
	Schengen     string        `csv:"Schengen"`
	EURO         string        `csv:"EURO"`
	UN           string        `csv:"UN"`
	UNmember     string        `csv:"UNmember"`
	Obsolete     string        `csv:"obsolete"`
	Cecilia2050  string        `csv:"Cecilia2050"`
	BRIC         string        `csv:"BRIC"`
	APEC         string        `csv:"APEC"`
	BASIC        string        `csv:"BASIC"`
	CIS          string        `csv:"CIS"`
	G7           string        `csv:"G7"`
	G20          string        `csv:"G20"`
	IEA          string        `csv:"IEA"`
	DACcode      string        `csv:"DACcode"`
}
```

### City data
```go
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
```

You can get more in depth view into our dataset by looking in `resources`

## Acknowledments
Key functionality was developed with the help of:

https://github.com/sams96/rgeo

https://github.com/IncSW/geoip2

https://github.com/dlclark/regexp2

https://www.geonames.org/

Library with extra functionality was inspired by:

https://pypi.org/project/country-converter/


## Road Map
- [ ] IP to city 
- [ ] Country to cities
- [ ] Performance / memory usage testing
- [ ] Full test coverage
- [ ] Better documentation
- [ ] Add more country details, eg. population
- [ ] Builds and releases workflow
- [ ] Installation guide
- [ ] Linting workflow

## License
Generally Apache v2.0, but
this product includes GeoLite2 data created by MaxMind, available from:

https://www.maxmind.com