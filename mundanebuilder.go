package builder

import (
	"github.com/Henry-Sarabia/blank"
	"github.com/pkg/errors"
	"io"
)

var (
	ErrSetterEmpty = errors.New("setters require a slice with at least 1 element")
	ErrNameBlank   = errors.New("setters require struct with valid 'Name' field")
)

// MundaneBuilder implements the Builder interface. MundaneBuilder randomly
// generates mundane Items from the provided resources.
type MundaneBuilder struct {
	Recipes    map[string]Recipe
	Attributes map[string]Attribute
	Groups     map[string]AttributeGroup
}

// NewMundaneBuilder returns an initialized MundaneBuilder.
func NewMundaneBuilder() MundaneBuilder {
	return MundaneBuilder{
		Recipes: make(map[string]Recipe),
		Attributes: make(map[string]Attribute),
		Groups: make(map[string]AttributeGroup),
	}
}

// SetRecipes reads the JSON-encoded Recipes from the provided Reader and
// adds them to the receiver's Recipe map.
func (mb *MundaneBuilder) SetRecipes(r io.Reader) error {
	rec, err := readRecipe(r)
	if err != nil {
		return errors.Wrap(err, "cannot read Recipes")
	}

	if err := mb.setRecipes(rec); err != nil {
		return errors.Wrap(err, "cannot set Recipes")
	}

	return nil
}

// setRecipes adds the provided Recipes to the receiver's Recipe map.
func (mb *MundaneBuilder) setRecipes(rec []Recipe) error {
	if len(rec) <= 0 {
		return ErrSetterEmpty
	}

	// TODO: Either implement duplicate checking or somehow allow duplicates. Latter is better.
	for _, v := range rec {
		if blank.Is(v.Name) {
			return ErrNameBlank
		}

		mb.Recipes[v.Name] = v
	}

	return nil
}

// SetAttributes reads the JSON-encoded Attributes from the provided Reader and
// adds them to the receiver's Attribute map.
func (mb *MundaneBuilder) SetAttributes(r io.Reader) error {
	attr, err := readAttribute(r)
	if err != nil {
		return errors.Wrap(err, "cannot read Attributes")
	}

	if err := mb.setAttributes(attr); err != nil {
		return errors.Wrap(err, "cannot set Attributes")
	}

	return nil
}

// setAttributes adds the provided Attributes to the receiver's Attribute map.
func (mb *MundaneBuilder) setAttributes(attr []Attribute) error {
	if len(attr) <= 0 {
		return ErrSetterEmpty
	}

	for _, v := range attr {
		if blank.Is(v.Name) {
			return ErrNameBlank
		}

		mb.Attributes[v.Name] = v
	}

	return nil
}

// SetAttributeGroups reads the JSON-encoded AttributeGroups from the provided
// Reader and adds them to the receiver's AttributeGroup map.
func (mb *MundaneBuilder) SetAttributeGroups(r io.Reader) error {
	grp, err := readAttributeGroup(r)
	if err != nil {
		return errors.Wrap(err, "cannot read AttributeGroups")
	}

	if err := mb.setAttributeGroups(grp); err != nil {
		return errors.Wrap(err, "cannot set AttributeGroups")
	}

	return nil
}

// setAttributeGroups adds the provided AttributeGroups to the receiver's
// AttributeGroup map.
func (mb *MundaneBuilder) setAttributeGroups(grp []AttributeGroup) error {
	if len(grp) <= 0 {
		return ErrSetterEmpty
	}

	for _, v := range grp {
		if blank.Is(v.Name) {
			return ErrNameBlank
		}

		mb.Groups[v.Name] = v
	}

	return nil
}

// LinkResources iterates through all of the receiver's maps and links each
// object's list of AttributeNames and AttributeGroupNames to their respective
// Attributes and AttributeGroups.
func (mb *MundaneBuilder) LinkResources() error {
	//TODO: Check to make sure the removal of slice of pointer data members didn't break linking code.
	if err := mb.linkGroups(); err != nil {
		return errors.Wrap(err, "cannot link groups") //TODO: Elaborate on error message
	}

	if err := mb.linkRecipes(); err != nil {
		return errors.Wrap(err, "cannot link recipes") //TODO: Elaborate
	}

	if err := mb.linkAffixes(); err != nil {
		return errors.Wrap(err, "cannot link affixes") //TODO: Elaborate
	}

	return nil
}

// linkGroups iterates through every AttributeGroup's AttributeNames and adds
// the corresponding Attribute addresses to the AttributeGroup's Attributes slice.
func (mb *MundaneBuilder) linkGroups() error {
	for _, grp := range mb.Groups {
		for _, name := range grp.AttributeNames {
			attr, ok := mb.Attributes[name]
			if !ok {
				return errors.Errorf("cannot find Attribute '%s' from AttributeGroup '%s' in builder's loaded Attributes", name, grp.Name)
			}

			grp.Attributes = append(grp.Attributes, &attr)
		}
	}
	return nil
}

// linkRecipes links every Recipe's AttributeNames and AttributeGroupNames to
// to their respective Attributes and AttributeGroups.
func (mb *MundaneBuilder) linkRecipes() error {
	for _, rec := range mb.Recipes {
		for _, comp := range rec.Comps {
			for _, prop := range comp.Properties {
				for _, name := range prop.AttributeNames {
					attr, ok := mb.Attributes[name]
					if !ok {
						return errors.Errorf("cannot find Attribute '%s' from Recipe '%s' in builder's loaded Attributes", name, rec.Name)
					}

					prop.Attributes = append(prop.Attributes, &attr)
				}

				for _, name := range prop.AttributeGroupNames {
					grp, ok := mb.Groups[name]
					if !ok {
						return errors.Errorf("cannot find AttributeGroup '%s' from Recipe '%s' in builder's loaded AttributeGroups", name, rec.Name)
					}

					prop.AttributeGroups = append(prop.AttributeGroups, &grp)
				}
			}
		}
	}
	return nil
}

// linkAffixes links every Attribute's AffixNames to their respective Affixes.
func (mb *MundaneBuilder) linkAffixes() error {
	for _, attr := range mb.Attributes {
		for _, name := range attr.PrefixNames {
			pfx, ok := mb.Attributes[name]
			if !ok {
				return errors.Errorf("cannot find Attribute '%s' from Attribute '%s' in builder's loaded Attributes", name, attr.Name)
			}

			attr.Prefixes = append(attr.Prefixes, &pfx)
		}
		for _, name := range attr.SuffixNames {
			sfx, ok := mb.Attributes[name]
			if !ok {
				return errors.Errorf("cannot find Attribute '%s' from Attribute '%s' in builder's loaded Attributes", name, attr.Name)
			}

			attr.Suffixes = append(attr.Suffixes, &sfx)
		}
	}

	return nil
}
