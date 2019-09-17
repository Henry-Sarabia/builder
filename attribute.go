package builder

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"log"
	"math/rand"
)

type Attribute struct {
	Name         string       `json:"name"`
	Common       VariantBlock `json:"common"`
	Uncommon     VariantBlock `json:"uncommon"`
	Rare         VariantBlock `json:"rare"`
	WeightFactor float64      `json:"weight_factor"`
	Prefixes     []string     `json:"prefix_references"`
}

func (a *Attribute) Reduce() Variant {
	c, u, r := len(a.Common.Variants), len(a.Uncommon.Variants), len(a.Rare.Variants)
	i := rand.Intn(c + u + r)

	var v Variant

	switch {
	case i < c:
		v = a.Common.Reduce()
	case i < c+u:
		v = a.Uncommon.Reduce()
	case i < c+u+r:
		v = a.Rare.Reduce()
	default:
		log.Fatal("Attribute.Reduce should never reach this") //TODO: Handle this case properly
	}

	v.WeightFactor = a.WeightFactor
	return v
}

// readAttribute reads the JSON-encoded Attributes from the provided Reader.
func readAttribute(r io.Reader) ([]Attribute, error) {
	var attr []Attribute

	if err := json.NewDecoder(r).Decode(&attr); err != nil {
		return nil, errors.Wrap(err, "cannot decode Attribute from io.Reader")
	}

	return attr, nil
}
