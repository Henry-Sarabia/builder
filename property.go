package builder

type Property struct {
	Name                string   `json:"name"`
	Required            bool     `json:"required"`
	AttributeNames      []string `json:"attribute_names"`
	AttributeGroupNames []string `json:"attribute_group_names"`
	Attributes          []*Attribute
	AttributeGroups     []*AttributeGroup
}

// Reduce returns a slice of randomly selected Attributes.
func (p *Property) Reduce() []Attribute {
	var attrs []Attribute
	for _, a := range p.Attributes {
		attrs = append(attrs, *a)
	}

	for _, grp := range p.AttributeGroups {
		for _, a := range grp.Attributes {
			attrs = append(attrs, *a)
		}
	}

	return attrs
}
