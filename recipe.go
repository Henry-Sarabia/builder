package builder

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"math/rand"
)

const ComponentSkipChance float64 = 0.5

type Recipe struct {
	Name       string      `json:"name"`
	BaseValue  float64     `json:"base_value"`
	BaseWeight float64     `json:"base_weight"`
	Comps      []Component `json:"components"`
}

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
