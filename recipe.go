package builder

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"math/rand"
)

const ComponentSkipChance float64 = 0.5

type Recipe struct {
	Name       string      `json:"name"`
	BaseValue  float64     `json:"base_value"`
	BaseWeight float64     `json:"base_weight"`
	Comps      []Component `json:"components"`
}

func (r *Recipe) Produce() Item {
	comps := r.Reduce()

	var atoms []AtomSlice
	for _, c := range comps {
		as := AtomSlice{
			ComponentLabel: c.Name,
			Atoms: c.Atomize(),
		}
		atoms = append(atoms, as)
	}

	val, wgt := r.BaseValue, r.BaseWeight
	for _, atms := range atoms {
		for _, a := range atms.Atoms {
			val *= a.ValueFactor
			wgt *= a.WeightFactor
		}
	}

	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err) // TODO: Handle this error better
	}

	return Item{
		ID: id.String(),
		Name: r.Name,
		Value: val,
		Weight: wgt,
		Description: stringify(atoms),
	}
}

func stringify(atoms []AtomSlice) string {
	var s string
	for _, atm := range atoms {
		s += atm.ComponentLabel
		for _, a := range atm.Atoms {
			s += a.String
		}
	}

	return s
}

// Reduce returns a slice of randomly selected Components.
func (r *Recipe) Reduce() []Component {
	var comps []Component
	for _, c := range r.Comps {
		if !c.Required && rand.Float64() <= ComponentSkipChance {
			continue
		}

		comps = append(comps, c)
	}

	if len(comps) <= 0 {
		i := rand.Intn(len(r.Comps))
		comps = append(comps, r.Comps[i])
	}

	return comps
}

// ReadRecipe reads the JSON-encoded Recipes from the provided Reader.
func readRecipe(r io.Reader) ([]Recipe, error) {
	var rec []Recipe

	if err := json.NewDecoder(r).Decode(&rec); err != nil {
		return nil, errors.Wrap(err, "cannot decode Recipe from io.Reader")
	}

	return rec, nil
}
