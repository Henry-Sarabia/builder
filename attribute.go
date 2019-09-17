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
	PrefixNames  []string     `json:"prefix_names"`
	SuffixNames  []string     `json:"suffix_names"`
	Prefixes     []*Attribute
	Suffixes     []*Attribute
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

	if pfx, ok := a.RandomPrefix(); ok {
		rd := pfx.Reduce()
		v.Name = v.Name + rd.Name
		v.ValueFactor = v.ValueFactor * rd.ValueFactor
	}

	if sfx, ok := a.RandomSuffix(); ok {
		rd := sfx.Reduce()
		v.Name = rd.Name + v.Name
		v.ValueFactor = v.ValueFactor * rd.ValueFactor
	}

	return v
}

func (a *Attribute) RandomPrefix() (*Attribute, bool) {
	if len(a.Prefixes) <= 0 {
		return nil, false
	}

	i := rand.Intn(len(a.Prefixes))
	return a.Prefixes[i], true
}

func (a *Attribute) RandomSuffix() (*Attribute, bool) {
	if len(a.Suffixes) <= 0 {
		return nil, false
	}

	i := rand.Intn(len(a.Suffixes))
	return a.Suffixes[i], true
}

// readAttribute reads the JSON-encoded Attributes from the provided Reader.
func readAttribute(r io.Reader) ([]Attribute, error) {
	var attr []Attribute

	if err := json.NewDecoder(r).Decode(&attr); err != nil {
		return nil, errors.Wrap(err, "cannot decode Attribute from io.Reader")
	}

	for _, a := range attr {
		if a.WeightFactor <= 0 {
			a.WeightFactor = 1
		}

		if a.Common.ValueFactor <= 0 {
			a.Common.ValueFactor = 1
		}

		if a.Uncommon.ValueFactor <= 0 {
			a.Uncommon.ValueFactor = 1
		}

		if a.Rare.ValueFactor <= 0 {
			a.Rare.ValueFactor = 1
		}
	}

	return attr, nil
}
