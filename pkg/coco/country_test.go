package coco_test

import (
	"testing"
)

func TestCountryNameToCountry(t *testing.T) {
	co, err := cc.CountryNameToCountry("Dem. People's Rep. of Korea", "")

	if err != nil {
		t.Error(err)
	}
	if co.ISO2 != "KP" {
		t.Error("Expected KP, got", co.ISO2)
	}
}

func TestCountryNameToCountryWithRegexInput(t *testing.T) {
	co, err := cc.CountryNameToCountry("France", "regex")

	if err != nil {
		t.Error(err)
	}
	if co.ISO2 != "FR" {
		t.Error("Expected FR, got", co.ISO2)
	}
}

func TestCountryNameToCountryWithISO2Input(t *testing.T) {
	co, err := cc.CountryNameToCountry("AL", "ISO2")

	if err != nil {
		t.Error(err)
	}
	if co.ISO2 != "AL" {
		t.Error("Expected FR, got", co.ISO2)
	}
}

func TestCountryNameToCountryWithISO3Input(t *testing.T) {
	co, err := cc.CountryNameToCountry("LTU", "ISO3")

	if err != nil {
		t.Error(err)
	}
	if co.ISO3 != "LTU" {
		t.Error("Expected LTU, got", co.ISO3)
	}
}

func TestCityNameToCountryISO2(t *testing.T) {
	ci, err := cc.CityNameToCity("Paryzius", true)

	if err != nil {
		t.Error(err)
	}
	if ci.CountryCode != "FR" {
		t.Error("Expected FR, got", ci.CountryCode)
	}
}

// Benchmark CountryNameToCountry
func BenchmarkCountryNameToCountryRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.CountryNameToCountry("Dem. People's Rep. of Korea", "")
	}
}

func BenchmarkCountryNameToCountryISO2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.CountryNameToCountry("KP", "ISO2")
	}
}
