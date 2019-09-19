package builder

import "math/rand"

const PropertySkipChance float64 = 0.5

type Component struct {
	Name       string     `json:"name"`
	Required   bool       `json:"required"`
	Properties []Property `json:"properties"`
}

// Reduce returns a slice of randomly selected Properties.
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

func (c *Component) Atomize() []Atom {
	var atoms []Atom

	props := c.Reduce()
	for _, p := range props {
		attr := p.Reduce()
		for _, a := range attr {
			v := a.Reduce()
			atom := Atom{
				PropertyLabel: p.Name,
				WeightFactor: v.WeightFactor,
				ValueFactor: v.ValueFactor,
				String: v.Name,
			}
			atoms = append(atoms, atom)
		}
	}

	return atoms
}