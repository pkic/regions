package regions

import (
	"sort"
	"strings"
)

// Country information
type Country struct {
	Name     string    `yaml:",omitempty"`
	ISO3166  string    `yaml:"iso3166-1,omitempty"` // ISO 3166-1 country code
	Complete bool      `yaml:""`
	Comments string    `yaml:",omitempty"`
	Regions  []*Region `yaml:",omitempty"`
}

// GetRegion returns a region by the code identifier (e.g., iso3361) and value
func (c *Country) GetRegion(regionName, codeName, codeValue string) *Region {
	for _, r := range c.Regions {
		if _, ok := r.Codes[codeName]; ok {
			if r.Codes[codeName] == codeValue {
				return r
			}
		}
	}

	// not ideal to search by name, but required as we do not always have an
	// common identifier
	if regionName != "" {
		for _, r := range c.Regions {
			for _, n := range r.Names {
				if strings.EqualFold(removeMetaData(regionName), n.Name) {
					return r
				}
			}
		}
	}

	return nil
}

// GetOrCreateRegion based on the code identifier (e.g., iso3361) and value
func (c *Country) GetOrCreateRegion(regionName []string, codeName, codeValue string) *Region {
	for _, rn := range regionName {
		r := c.GetRegion(rn, codeName, codeValue)
		if r != nil {
			r.Codes[codeName] = codeValue
			return r
		}
	}

	c.Regions = append(c.Regions, &Region{
		Codes: map[string]string{codeName: codeValue},
	})

	c.sortRegions()
	return c.GetRegion("", codeName, codeValue)
}

// RemoveSource removes data attributed to a source
func (c *Country) RemoveSource(source string) error {
	for _, r := range c.Regions {
		err := r.RemoveSource(source)
		if err != nil {
			return err
		}
	}

	// remove regions without any data left
	for i := 0; i < len(c.Regions); i++ {
		if len(c.Regions[i].Names) == 0 {
			c.Regions = append(c.Regions[:i], c.Regions[i+1:]...)
			i--
		}
	}

	return nil
}

func (c *Country) sortRegions() {
	// Sort by ISO3166Code to preserver order in updates
	sort.Slice(c.Regions, func(i, j int) bool {
		if len(c.Regions[i].Names) == 0 || len(c.Regions[j].Names) == 0 {
			return true
		}
		switch strings.Compare(c.Regions[i].Names[0].Name, c.Regions[j].Names[0].Name) {
		case -1:
			return true
		case 1:
			return false
		}
		return c.Regions[i].Names[0].Name > c.Regions[j].Names[0].Name
	})
}

// MarshalYAML is a custom mashaller to sort and restoring comments automatically
func (c *Country) MarshalYAML() (interface{}, error) {
	c.sortRegions()

	type alias *Country
	// node := yaml.Node{}
	// err := node.Encode(alias(c))
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding county head comment in %q: %w", c.ISO3166, err)
	// }

	// node.HeadComment = "File automatically udpated!\n"
	// node.HeadComment += " - Only update non automated sources; or your changes will get lost.\n"
	// node.HeadComment += " - Comments are currently not preserved in automated updates.\n"
	// node.HeadComment += "\n"

	// return node, nil
	return alias(c), nil
}
