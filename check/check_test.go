package check

import (
	"testing"
)

func TestIsCountryRegion(t *testing.T) {
	country := "US"
	region := "Alabama"
	expected := error(nil)
	actual := IsCountryRegion(country, region)
	if expected != actual {
		t.Error(
			"For region", region,
			"in country", country,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsCountryRegionUppercase(t *testing.T) {
	country := "US"
	region := "ALABAMA"
	expected := error(nil)
	actual := IsCountryRegion(country, region)
	if expected != actual {
		t.Error(
			"For region", region,
			"in country", country,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsCountryRegionISOCode(t *testing.T) {
	country := "US"
	region := "US-AL"
	expected := error(nil)
	actual := IsCountryRegion(country, region)
	if expected != actual {
		t.Error(
			"For region", region,
			"in country", country,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsCountryRegionDiacritics(t *testing.T) {
	country := "FR"
	region := "Rhône"
	expected := error(nil)
	actual := IsCountryRegion(country, region)
	if expected != actual {
		t.Error(
			"For region", region,
			"in country", country,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsCountryRegionDiacriticsMissing(t *testing.T) {
	country := "FR"
	region := "Rhone"
	expected := error(nil)
	actual := IsCountryRegion(country, region)
	if expected != actual {
		t.Error(
			"For region", region,
			"in country", country,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsCountryRegionNormalized(t *testing.T) {
	country := "GB"
	region := "London"
	expected := error(nil)
	actual := IsCountryRegion(country, region)
	if expected != actual {
		t.Error(
			"For region", region,
			"in country", country,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsCountryRegionNotNormalized(t *testing.T) {
	country := "GB"
	region := "London, City of"
	expected := error(nil)
	actual := IsCountryRegion(country, region)
	if expected != actual {
		t.Error(
			"For region", region,
			"in country", country,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsCountryRegionMissingHyphenIgnored(t *testing.T) {
	country := "NL"
	region := "Noord Holland"
	expected := error(nil)
	actual := IsCountryRegion(country, region)
	if expected != actual {
		t.Error(
			"For region", region,
			"in country", country,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsCountryRegionErrCountryNotExist(t *testing.T) {
	country := "USX"
	region := "US-AL"
	expected := ErrCountryNotExist
	actual := IsCountryRegion(country, region)
	if expected != actual {
		t.Error(
			"For region", region,
			"in country", country,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsCountryRegionErrRegionUnknown(t *testing.T) {
	country := "US"
	region := "Does not exist"
	expected := ErrRegionUnknown
	actual := IsCountryRegion(country, region)
	if expected != actual {
		t.Error(
			"For region", region,
			"in country", country,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsCountryRegionErrRegionNotExist(t *testing.T) {
	country := "NL"
	region := "Does not exist"
	expected := ErrRegionNotExist
	actual := IsCountryRegion(country, region)
	if expected != actual {
		t.Error(
			"For region", region,
			"in country", country,
			"expected", expected,
			"got", actual,
		)
	}
}

func BenchmarkIsCountryRegion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsCountryRegion("US", "US-AL")
	}
}

func BenchmarkIsCountryRegionDiacritics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsCountryRegion("FR", "Rhône")
	}
}
