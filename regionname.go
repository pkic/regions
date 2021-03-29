package regions

import (
	"sort"

	yaml "gopkg.in/yaml.v3"
)

// RegionName holds the languages and sources of a region name
type RegionName struct {
	Name      string   `yaml:",omitempty"`
	Languages []string `yaml:",omitempty,flow"`
	Sources   []string `yaml:",omitempty,flow"`
}

// MarshalYAML is a custom mashaller to sort and restoring comments automatically
func (n *RegionName) MarshalYAML() (interface{}, error) {
	sort.Strings(n.Languages)
	sort.Strings(n.Sources)

	type alias *RegionName
	node := yaml.Node{}
	err := node.Encode(alias(n))
	if err != nil {
		return nil, err
	}

	node.HeadComment = "\n\n --"

	return node, nil
}
