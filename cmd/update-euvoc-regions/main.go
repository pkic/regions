package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/knakk/sparql"
	yaml "gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("missing arguments, usage: `%s ./output-directory`\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	var err error
	e := &Europe{}

	e.client, err = sparql.NewRepo("http://publications.europa.eu/webapi/rdf/sparql", sparql.Timeout(5*time.Minute))
	if err != nil {
		log.Fatal(err)
	}

	countries, err := e.getCountries()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range countries {
		countryFile := fmt.Sprintf("%s/%s.yaml", os.Args[1], strings.ToLower(c.ISO3166))

		// Preserve existing data
		currentData, err := ioutil.ReadFile(countryFile)
		if err == nil {
			err = yaml.Unmarshal(currentData, &c)
			if err != nil {
				log.Fatal(err)
			}
		}

		err = e.getRegions(&c)
		if err != nil {
			log.Fatal(err)
		}

		updatedData, err := yaml.Marshal(c)
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
