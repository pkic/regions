package regions

import (
	"sort"

	yaml "gopkg.in/yaml.v3"
)

// Source holds the original value and language
type Source struct {
	Name      string   `yaml:",omitempty"`
	Languages []string `yaml:",omitempty,flow"`
	Value     string   `yaml:",omitempty"` // Value as included in the source
	Type      string   `yaml:",omitempty"` // To indicate alternative or other non primary names
	//Status    string    `yaml:",omitempty"`
	//EndDate   time.Time `yaml:",omitempty"`
}

// MarshalYAML is a custom mashaller to sort and restoring comments automatically
func (s *Source) MarshalYAML() (interface{}, error) {
	sort.Strings(s.Languages)

	type alias *Source
	node := yaml.Node{}
	err := node.Encode(alias(s))
	if err != nil {
		return nil, err
	}

	return node, nil
}
