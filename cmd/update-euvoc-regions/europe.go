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
func (e *Europe) getRegions(c *regions.Country) error {
	res, err := e.client.Query(`
	  PREFIX skos: <http://www.w3.org/2004/02/skos/core#>
	  PREFIX euvoc: <http://publications.europa.eu/ontology/euvoc#>
	  PREFIX status: <http://publications.europa.eu/resource/authority/concept-status/>
	  PREFIX ev: <http://eurovoc.europa.eu/>
	  PREFIX country: <http://publications.europa.eu/resource/authority/country>
	  SELECT DISTINCT ?subRegion, ?subRegionLabel
	  WHERE {  
		  ?subRegion skos:prefLabel ?subRegionLabel ; 
					  euvoc:status ?subRegionStatus ;
					  skos:inScheme ev:100278 ;
					  skos:broader+ ?regionGroup . 
					  FILTER(?subRegionStatus = status:CURRENT) .
 
		  OPTIONAL {
			  ?subRegion skos:altLabel ?subAltRegionLabel FILTER(langMatches(lang(?subAltRegionLabel), "en")) .
		  }
 
		  OPTIONAL {
			  ?regionGroup skos:prefLabel ?regionGroupLabel ;
					  skos:topConceptOf ev:100278 .
					  FILTER(langMatches(lang(?regionGroupLabel), "en")) .  
		  }
	  
		  ?country skos:prefLabel ?countryLabel ;  
					  skos:related ?regionGroup ; 
					  skos:inScheme  ev:100277 .
					  FILTER(langMatches(lang(?countryLabel), "en") AND str(?countryLabel) = "` + c.Name + `") . 
	  }`)
	if err != nil {
		return err
	}

	err = c.RemoveSource(sourceIdentifier)
	if err != nil {
		return err
	}

	for _, v := range res.Results.Bindings {
		r := c.GetOrCreateRegion([]string{v["subRegionLabel"].Value}, sourceIdentifier, v["subRegion"].Value)
		err = r.Add(v["subRegionLabel"].Value, strings.ToLower(v["subRegionLabel"].Lang), sourceIdentifier)
		if err != nil {
			return err
		}
	}

	return nil
}
