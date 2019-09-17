package builder

import "math/rand"

type VariantBlock struct {
	ValueFactor float64  `json:"value_factor"`
	Variants    []string `json:"variants"`
}

type Variant struct {
	Name         string
	ValueFactor  float64
	WeightFactor float64
}

// Reduce returns a randomly selected Variant.
func (vb *VariantBlock) Reduce() Variant {
	r := rand.Intn(len(vb.Variants))

	return Variant{
		Name:        vb.Variants[r],
		ValueFactor: vb.ValueFactor,
	}
}
