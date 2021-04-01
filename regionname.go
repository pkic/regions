package regions

import (
	"sort"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

// RegionName holds the languages and sources of a region name
type RegionName struct {
	Name    string    `yaml:",omitempty"` // Normalized value
	Sources []*Source `yaml:",omitempty"`
}

// Remove source attribution
func (r *RegionName) removeSource(source string) error {
	for sid := 0; sid < len(r.Sources); sid++ {
		if strings.HasPrefix(strings.ToLower(r.Sources[sid].Name), strings.ToLower(source)) {
			if len(r.Sources) == 1 {
				r.Sources = nil
				return nil
			}
			r.Sources = append(r.Sources[:sid], r.Sources[sid+1:]...)
			sid--
		}
	}
	return nil
}

// Add source attribution including value at source and language
func (n *RegionName) addSource(normalizedName, regionName, language, source string) error {
	for _, s := range n.Sources {
		if strings.EqualFold(s.Name, source) &&
			(s.Value == "" || strings.EqualFold(s.Value, regionName)) {

			if !stringInSlice(s.Languages, language) && language != "" {
				s.Languages = append(s.Languages, strings.ToLower(language))
			}
			return nil
		}
	}

	s := &Source{Name: source}
	if regionName != normalizedName {
		s.Value = regionName
	}
	if language != "" {
		s.Languages = []string{strings.ToLower(language)}
	}
	n.Sources = append(n.Sources, s)
	return nil
}

func (n *RegionName) sortSources() {
	sort.Slice(n.Sources, func(i, j int) bool {
		switch strings.Compare(n.Sources[i].Name, n.Sources[j].Name) {
		case -1:
			return true
		case 1:
			return false
		}
		return n.Sources[i].Name > n.Sources[j].Name
	})
}

// MarshalYAML is a custom mashaller to sort and restoring comments automatically
func (n *RegionName) MarshalYAML() (interface{}, error) {
	n.sortSources()

	type alias *RegionName
	node := yaml.Node{}
	err := node.Encode(alias(n))
	if err != nil {
		return nil, err
	}

	return node, nil
}
