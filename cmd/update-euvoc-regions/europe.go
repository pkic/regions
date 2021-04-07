package main

import (
	"strings"

	"github.com/pkic/regions"

	"github.com/knakk/sparql"
)

var (
	sourceIdentifier = "euvoc"
)

type Europe struct {
	client *sparql.Repo
}

// getCountries requests the countries in Europe including the ISO 32166-1 country code
func (e *Europe) getCountries() ([]regions.Country, error) {
	res, err := e.client.Query(`
	 PREFIX skos: <http://www.w3.org/2004/02/skos/core#>
	 PREFIX euvoc: <http://publications.europa.eu/ontology/euvoc#>
	 PREFIX status: <http://publications.europa.eu/resource/authority/concept-status/>
	 PREFIX dc: <http://purl.org/dc/terms/>
	 PREFIX continent: <http://publications.europa.eu/resource/authority/continent/>
	 PREFIX ogcgs: <http://www.opengis.net/ont/geosparql#>
	 PREFIX nt: <http://publications.europa.eu/resource/authority/notation-type/>
	 SELECT DISTINCT ?countryCode ?countryLabel
	 FROM <http://publications.europa.eu/resource/authority/country>
	 FROM <http://publications.europa.eu/resource/cellar/9f2bd600-ae7b-11e7-837e-01aa75ed71a1>
	 WHERE {
		 ?countryURI skos:prefLabel ?countryLabel ;
					 ogcgs:sfWithin continent:EUROPE ;
					 euvoc:status status:CURRENT .
	 
		  ?x rdfs:label ?countryLabel .
		  ?countryURI euvoc:xlNotation  ?xlNotation  .
		  ?xlNotation  dc:type ?type .
		  ?xlNotation euvoc:xlCodification ?countryCode .
		  FILTER(LANG(?countryLabel)="en" AND ?type = nt:ISO_3166_1_ALPHA_2) .
	 }`)
	if err != nil {
		return nil, err
	}

	var result []regions.Country
	for _, v := range res.Results.Bindings {
		result = append(result, regions.Country{
			Name:    v["countryLabel"].Value,
			ISO3166: strings.ToUpper(v["countryCode"].Value),
		})
	}

	return result, nil
}

// getRegions retrieves and updates the region data
// TODO: Include information about euvoc:status, euvoc:endDate and dc:isReplacedBy
func (e *Europe) getRegions(c *regions.Country) error {
	res, err := e.client.Query(`
	PREFIX skos: <http://www.w3.org/2004/02/skos/core#>
	PREFIX euvoc: <http://publications.europa.eu/ontology/euvoc#>
	PREFIX status: <http://publications.europa.eu/resource/authority/concept-status/>
	PREFIX ev: <http://eurovoc.europa.eu/>
	PREFIX country: <http://publications.europa.eu/resource/authority/country>
	SELECT DISTINCT ?region, ?regionLabel, ?altRegionLabel
	WHERE {
	  ?regionsOf skos:prefLabel ?regionsOfLabel;
	      euvoc:status status:CURRENT ;
		  skos:inScheme ev:100141 .
		  FILTER(STRSTARTS(?regionsOfLabel, "regions ")) .  
		  FILTER(STRENDS(?regionsOfLabel, " ` + c.Name + `")) .  

	   ?region skos:prefLabel ?regionLabel ;
	      euvoc:status status:CURRENT ;
		  skos:broader+ ?regionsOf .

	  OPTIONAL {
		  ?region skos:altLabel ?altRegionLabel ;
			euvoc:status status:CURRENT .
			FILTER(langMatches(lang(?altRegionLabel), "en")) .
	  }
	}`)
	if err != nil {
		return err
	}

	err = c.RemoveSource(sourceIdentifier)
	if err != nil {
		return err
	}

	for _, v := range res.Results.Bindings {
		r := c.GetOrCreateRegion([]string{v["regionLabel"].Value, v["altRegionLabel"].Value}, sourceIdentifier, v["region"].Value)

		err = r.Add(v["regionLabel"].Value, strings.ToLower(v["regionLabel"].Lang), sourceIdentifier, "")
		if err != nil {
			return err
		}

		if v["altRegionLabel"].Value != "" {
			err = r.Add(v["altRegionLabel"].Value, strings.ToLower(v["altRegionLabel"].Lang), sourceIdentifier, "alternative")
			if err != nil {
				return err
			}
		}
	}

	return nil
}
