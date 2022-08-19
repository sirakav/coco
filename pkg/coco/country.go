package coco

import (
	"errors"
	"log"
	"strconv"

	"github.com/cornelk/hashmap"
	"github.com/dlclark/regexp2"
	"github.com/gocarina/gocsv"

	"github.com/sirakav/coco/resources"
	"strings"
)

type CompiledRegex struct {
	*regexp2.Regexp
}

func (cr *CompiledRegex) MarshalCSV() (string, error) {
	return "", nil
}

func (cr *CompiledRegex) UnmarshalCSV(csv string) (err error) {
	cr.Regexp = regexp2.MustCompile(csv, regexp2.IgnoreCase)
	return nil
}

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

func loadCountryData() []*CountryData {
	c := []*CountryData{}
	if err := gocsv.UnmarshalBytes(resources.CountryDataCSV, &c); err != nil {
		panic(err)
	}

	return c
}

func getISO2ToCountryHM(c []*CountryData) *hashmap.HashMap[string, *CountryData] {
	hm := hashmap.New[string, *CountryData]()
	for _, co := range c {
		hm.Set(strings.ToLower(co.ISO2), co)
	}
	return hm
}

func getISO3ToCountryHM(c []*CountryData) *hashmap.HashMap[string, *CountryData] {
	hm := hashmap.New[string, *CountryData]()
	for _, co := range c {
		hm.Set(strings.ToLower(co.ISO3), co)
	}
	return hm
}

func getISONumericToCountryHM(c []*CountryData) *hashmap.HashMap[string, *CountryData] {
	hm := hashmap.New[string, *CountryData]()
	for _, co := range c {
		hm.Set(co.ISOnumeric, co)
	}
	return hm
}

func getCountryNameInputFormat(name string) string {
	_, err := strconv.Atoi(name)
	if err == nil {
		return "ISOnumeric"
	}

	if len(name) == 2 {
		return "ISO2"
	} else if len(name) == 3 {
		return "ISO3"
	}
	return "regex"
}

func (c *COCO) matchCountryRegex(name string) (*CountryData, error) {
	for _, c := range c.Countries {
		m, err := c.Regex.MatchString(name)
		if err != nil {
			log.Fatal(err)
		}

		if m {
			return c, nil
		}
	}
	return &CountryData{}, &CountryNotFoundError{}
}

func getCountryFromHashMap(h *hashmap.HashMap[string, *CountryData], name string) (*CountryData, error) {
	v, ok := h.Get(name)
	if !ok {
		return &CountryData{}, &CountryNotFoundError{}
	}
	return v, nil
}

func (c *COCO) matchISO2ToCountry(name string) (*CountryData, error) {
	return getCountryFromHashMap(c.iso2ToCountryHM, strings.ToLower(name))
}

func (c *COCO) matchISO3ToCountry(name string) (*CountryData, error) {
	return getCountryFromHashMap(c.iso3ToCountryHM, strings.ToLower(name))
}

func (c *COCO) matchISONumericToCountry(name string) (*CountryData, error) {
	return getCountryFromHashMap(c.isoNumericToCountryHM, strings.ToLower(name))
}

func (c *COCO) getCountryData(name string, inputFormat string) (*CountryData, error) {
	if inputFormat == "" {
		inputFormat = getCountryNameInputFormat(name)
	}
	switch inputFormat {
	case "regex":
		return c.matchCountryRegex(name)
	case "ISO2":
		return c.matchISO2ToCountry(name)
	case "ISO3":
		return c.matchISO3ToCountry(name)
	case "ISOnumeric":
		return c.matchISONumericToCountry(name)
	}

	return &CountryData{}, errors.New("input format not supported")
}

func (c *COCO) CountryNameToCountry(name string, inputFormat string) (*CountryData, error) {
	co, err := c.getCountryData(name, inputFormat)
	if err != nil {
		return &CountryData{}, err
	}
	return co, nil
}

func (c *COCO) CountryNameToISO2(name string, inputFormat string) (string, error) {
	co, err := c.getCountryData(name, inputFormat)
	return co.ISO2, err
}

func (c *COCO) CountryNameToISO3(name string, inputFormat string) (string, error) {
	co, err := c.getCountryData(name, inputFormat)
	return co.ISO3, err
}

func (c *COCO) CountryNameToISONumeric(name string, inputFormat string) (string, error) {
	co, err := c.getCountryData(name, inputFormat)
	return co.ISOnumeric, err
}

func (c *COCO) CityNameToCountryISO2(name string) (string, error) {
	ci, err := c.CityNameToCity(name, true)

	if err != nil {
		return "", err
	}
	return ci.CountryCode, nil
}

func (c *COCO) CityNameToCountryISO3(name string) (string, error) {
	ci, err := c.CityNameToCity(name, true)

	if err != nil {
		return "", err
	}
	return ci.CountryCode, nil
}

func (c *COCO) CityNameToCountry(name string) (*CountryData, error) {
	ci, err := c.CityNameToCity(name, true)
	if err != nil {
		return &CountryData{}, err
	}

	co, err := c.matchISO2ToCountry(ci.CountryCode)
	if err != nil {
		return &CountryData{}, err
	}
	return co, nil
}
