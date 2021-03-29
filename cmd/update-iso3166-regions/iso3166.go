package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pkic/regions"
)

var (
	sourceIdentifier = "iso3166-2"
)

// ISO3166 implementation
type ISO3166 struct {
	records [][]string
}

func (iso *ISO3166) sort() {
	// Sort by ISO3166Code to preserver order in updates
	sort.Slice(iso.records, func(i, j int) bool {
		switch strings.Compare(iso.records[i][4], iso.records[j][4]) {
		case -1:
			return true
		case 1:
			return false
		}
		return iso.records[i][4] > iso.records[j][4]
	})
}

// getData fetches the response body bytes from an HTTP get to the provider url,
// or returns an error.
func (iso *ISO3166) getData(url string) error {
	data, err := ioutil.ReadFile(url)
	if err != nil {
		return err
	}

	if strings.ToLower(filepath.Ext(url)) == ".zip" {
		zipReader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
		if err != nil {
			return err
		}

		for _, zipFile := range zipReader.File {
			if strings.ToLower(filepath.Ext(zipFile.Name)) != ".csv" {
				continue
			}

			csvReader, err := zipFile.Open()
			if err != nil {
				return err
			}
			defer csvReader.Close()

			data, err = ioutil.ReadAll(csvReader)
			if err != nil {
				return err
			}

			// we only read the first csv file
			break
		}
	}

	csvReader := csv.NewReader(bytes.NewReader(data))
	csvReader.TrimLeadingSpace = true
	iso.records, err = csvReader.ReadAll()
	if err != nil {
		return err
	}

	iso.sort()

	return nil
}

func (iso *ISO3166) getCountries() ([]regions.Country, error) {
	// filter duplicates
	c := make(map[string]bool)
	for _, record := range iso.records {
		c[strings.ToUpper(record[0])] = true
	}

	// prepare result
	var result []regions.Country
	for cc := range c {
		if len(cc) == 2 {
			result = append(result, regions.Country{
				ISO3166: cc,
			})
		}
	}
	return result, nil
}

func (iso *ISO3166) getRegions(c *regions.Country) error {
	err := c.RemoveSource(sourceIdentifier)
	if err != nil {
		return err
	}

	for _, record := range iso.records {
		if !strings.EqualFold(c.ISO3166, record[0]) {
			continue
		}

		// ISO 3166-2 CSV file
		regionType := record[3]
		regionCode := strings.ToUpper(record[4])
		regionLang := strings.ToLower(record[5])
		regionName := filterValue(record[7])
		regionNameLocal := filterValue(record[8])

		// other region types such as country, arctic region, etc
		// TODO: filter all region types
		if regionType == "398" || regionType == "413" || regionType == "345" {
			continue
		}

		r := c.GetOrCreateRegion([]string{regionName, regionNameLocal}, sourceIdentifier, regionCode)
		err = r.Add(regionName, regionLang, sourceIdentifier)
		if err != nil {
			return err
		}
		if regionNameLocal != "" {
			err = r.Add(regionNameLocal, regionLang, sourceIdentifier)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func filterValue(value string) string {
	if strings.ContainsAny(value, ":=/_") {
		return ""
	}

	return value
}
