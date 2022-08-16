package coco_test

import (
	"testing"
)

func TestIPToISO2(t *testing.T) {
	s, err := cc.IPToISO2("18.1.154.3")

	if err != nil {
		t.Error(err)
	}
	if s != "US" {
		t.Error("Expected US, got", s)
	}
}

func TestIPToISO3(t *testing.T) {
	s, err := cc.IPToISO3("59.1.54.3")

	if err != nil {
		t.Error(err)
	}
	if s != "KOR" {
		t.Error("Expected KOR, got", s)
	}
}

func TestIPToCountry(t *testing.T) {
	s, err := cc.IPToCountry("59.1.54.3")
	if err != nil {
		t.Error(err)
	}
	if s.NameShort != "South Korea" {
		t.Error("Expected South Korea, got", s.NameShort)
	}
}
