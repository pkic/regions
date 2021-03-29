package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("missing arguments, usage: `%s iso3166-2.csv ./output-directory`\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	iso := ISO3166{}
	// Get all ISO 3166-2 regions
	err := iso.getData(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	countries, err := iso.getCountries()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range countries {
		countryFile := fmt.Sprintf("%s/%s.yaml", os.Args[2], strings.ToLower(c.ISO3166))

		// Preserve existing data
		currentData, err := ioutil.ReadFile(countryFile)
		if err == nil {
			err = yaml.Unmarshal(currentData, &c)
			if err != nil {
				log.Fatal(err)
			}
		}

		err = iso.getRegions(&c)
		if err != nil {
			log.Fatal(err)
		}

		updatedData, err := yaml.Marshal(&c)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(countryFile, updatedData, 0644)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Finished %s (%s) with %d regions", c.Name, c.ISO3166, len(c.Regions))
	}
}
