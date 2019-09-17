package builder

import "math/rand"

const PropertySkipChance float64 = 0.5

type Component struct {
	Name       string     `json:"name"`
	Required   bool       `json:"required"`
	Properties []Property `json:"properties"`
}

func (c *Component) Reduce() []Property {
	var props []Property
	for _, p := range c.Properties {
		if !p.Required && rand.Float64() <= PropertySkipChance {
			continue
		}

		props = append(props, p)
	}

	if len(props) <= 0 {
		i := rand.Intn(len(c.Properties))
		props = append(props, c.Properties[i])
	}

	return props
}
