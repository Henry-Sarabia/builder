package builder

type Atom struct {
	PropertyLabel string
	WeightFactor  float64
	ValueFactor   float64
	String        string
}

type AtomSlice struct {
	ComponentLabel string
	Atoms []Atom
}