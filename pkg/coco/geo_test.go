package coco_test

import (
	"testing"
)

func TestLatLonToISO2(t *testing.T) {
	s, err := cc.LatLonToISO2(25.2663299, 54.6977051)

	if err != nil {
		t.Error(err)
	}
	if s != "LT" {
		t.Error("Expected LT, got", s)
	}
}

func TestLatLonToISO3(t *testing.T) {
	s, err := cc.LatLonToISO3(60.1551229, 37.2166344)

	if err != nil {
		t.Error(err)
	}
	if s != "TKM" {
		t.Error("Expected TKM, got", s)
	}
}

func TestLatLonToCountry(t *testing.T) {
	s, err := cc.LatLonToCountry(-74.456272,-9.8826415)
	if err != nil {
		t.Error(err)
	}
	if s.NameShort != "Peru" {
		t.Error("Expected Peru, got", s.NameShort)
	}
}

func BenchmarkLatLonToISO3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cc.LatLonToISO3(25.2663299, 54.6977051)
	}
}
