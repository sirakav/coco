package coco

type CountryNotFoundError struct{}

func (e *CountryNotFoundError) Error() string {
	return "country not found"
}

type CityNotFoundError struct{}

func (e *CityNotFoundError) Error() string {
	return "country not found"
}
