package builder

import "math/rand"

type VariantBlock struct {
	ValueFactor factor   `json:"value_factor"`
	Variants    []string `json:"variants"`
}

type Variant struct {
	Name         string
	ValueFactor  factor
	WeightFactor factor
}

func (vb *VariantBlock) Reduce() Variant {
	r := rand.Intn(len(vb.Variants))

	return Variant{
		Name:        vb.Variants[r],
		ValueFactor: vb.ValueFactor,
	}
}
