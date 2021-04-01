package regions

import (
	"fmt"
	"sort"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

// Region hold the information about a region
type Region struct {
	Codes map[string]string `yaml:",omitempty"` // ISO 3166-2 region code, EU identifier, etc
	Names []*RegionName     `yaml:",omitempty"`
}

// RemoveSource removes regions attributed to a source
//
// Matching is done based on a prefix, this allows to remove all sources from a
// given website.
//
// Language attributions are not removed, when the region exists in other sources
//
// When a region is attributed to multiple sources the region is retained and only
// the attribution to the removed source is removed.
func (r *Region) RemoveSource(source string) error {
	delete(r.Codes, source)

	for nid := 0; nid < len(r.Names); nid++ {
		err := r.Names[nid].removeSource(source)
		if err != nil {
			return err
		}

		// Remove region name if no sources left
		if r.Names[nid].Sources == nil || len(r.Names[nid].Sources) == 0 {
			r.Names = append(r.Names[:nid], r.Names[nid+1:]...)
			nid--
		}
	}

	return nil
}

// Add a new region
func (r *Region) Add(regionName, language, source string) error {
	normalizedName := removeMetaData(regionName)
	if normalizedName == "" {
		normalizedName = regionName
	}

	for _, n := range r.Names {
		// If exists, check if source is listed, else add for reference
		if strings.EqualFold(n.Name, normalizedName) {
			return n.addSource(normalizedName, regionName, language, source)
		}
	}

	// Add name if not present
	rn := &RegionName{Name: normalizedName}
	err := rn.addSource(normalizedName, regionName, language, source)
	if err != nil {
		return err
	}

	r.Names = append(r.Names, rn)

	return nil
}

func (r *Region) sort() {
	// Sort by name to preserver order in updates
	sort.Slice(r.Names, func(i, j int) bool {
		switch strings.Compare(r.Names[i].Name, r.Names[j].Name) {
		case -1:
			return true
		case 1:
			return false
		}
		return r.Names[i].Name > r.Names[j].Name
	})
}

// String returns the region name in English or the first name when no English
// name is present.
func (r *Region) String() string {
	for _, n := range r.Names {
		for _, s := range n.Sources {
			for _, l := range s.Languages {
				if l == "en" {
					return n.Name
				}
			}
		}
	}

	if len(r.Names) > 0 {
		return r.Names[0].Name
	}
	return ""
}

// MarshalYAML is a custom mashaller to sort and restoring comments automatically
func (r *Region) MarshalYAML() (interface{}, error) {
	r.sort()

	type alias *Region
	node := yaml.Node{}
	err := node.Encode(alias(r))
	if err != nil {
		return nil, err
	}

	node.HeadComment = fmt.Sprintf("\n\n%s", r.String())

	return node, nil
}
