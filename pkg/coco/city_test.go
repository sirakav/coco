package coco_test

import (
	"testing"
)

func TestCityNameToCity(t *testing.T) {
	ci, err := cc.CityNameToCity("PArIs", false)

	if err != nil {
		t.Error(err)
	}
	if ci.CountryCode != "FR" {
		t.Error("Expected FR, got", ci.CountryCode)
	}
}

func TestCityNameToCityWithExtendedSearch(t *testing.T) {
	ci, err := cc.CityNameToCity("LondoNas", true)

	if err != nil {
		t.Error(err)
	}
	if ci.CountryCode != "GB" {
		t.Error("Expected GB, got", ci.CountryCode)
	}

	ci, err = cc.CityNameToCity("پارىژ", true)

	if err != nil {
		t.Error(err)
	}
	if ci.CountryCode != "FR" {
		t.Error("Expected FR, got", ci.CountryCode)
	}
}

func BenchmarkCityNameToCityWithExtendedSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.CityNameToCity("Paryzius", true)
	}
}

func BenchmarkCityNameToCity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.CityNameToCity("paris", false)
	}
}
