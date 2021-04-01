package check

import (
	"errors"
	"strings"
)

var (
	ErrCountryNotExist = errors.New("country does not exist")
	ErrRegionUnknown   = errors.New("region unknown")
	ErrRegionNotExist  = errors.New("region does not exist")
)

// This package uses the `create-cc-map` command to generate a `regionMap` map.
//go:generate create-cc-map ../data/ ./regions_map.go

// IsCountryRegion returns true if the input is a known region for
// the given country.
func IsCountryRegion(country, region string) error {
	country = strings.ToUpper(country)
	if c, ok := regionMap[country]; ok {
		if !(len(region) == 5 && strings.HasPrefix(region, country+"-")) {
			region = strings.Replace(region, "-", " ", -1)
		}
		if _, ok := c.regions[strings.ToUpper(region)]; ok {
			return nil
		}
		// Region not defined in country
		if c.complete {
			// Region data is indicated as complete
			return ErrRegionNotExist
		}
		// Region data might be incomplete
		return ErrRegionUnknown
	}
	// Unkown country
	return ErrCountryNotExist
}
