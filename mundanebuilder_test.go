package builder

import (
	"github.com/davecgh/go-spew/spew"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

const (
	wildcard = "/*.json"
	testFileRecipe = "testdata/recipes/art.json"
	testFileGroup = "testdata/groups/base.json"
	testDirRecipe = "testdata/recipes"
	testDirAttribute = "testdata/attributes"
	testDirAttributeGroup = "testdata/groups"
)

func setupMundaneBuilder() (*MundaneBuilder, error) {
	b := NewMundaneBuilder()

	rec, err := os.Open(testFileRecipe)
	if err != nil {
		return nil, err
	}
	if err := b.SetRecipes(rec); err != nil {
		return nil, err
	}

	grp, err := os.Open(testFileGroup)
	if err != nil {
		return nil, err
	}
	if err := b.SetAttributeGroups(grp); err != nil {
		return nil, err
	}

	attr, err := filepath.Glob(testDirAttribute + wildcard)
	if err != nil {
		return nil, err
	}

	for _, a := range attr {
		f, err := os.Open(a)
		if err != nil {
			return nil, err
		}

		if err := b.SetAttributes(f); err != nil {
			return nil, err
		}
	}

	return &b, nil
}

func TestMundaneBuilder_Item(t *testing.T) {
	rand.Seed(1)

	b := NewMundaneBuilder()

	rec, err := filepath.Glob(testDirRecipe + wildcard)
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range rec {
		f, err := os.Open(r)
		if err != nil {
			t.Fatal(err)
		}

		if err := b.SetRecipes(f); err != nil {
			t.Fatal(err)
		}
	}

	attr, err := filepath.Glob(testDirAttribute + wildcard)
	if err != nil {
		t.Fatal(err)
	}

	for _, a := range attr {
		f, err := os.Open(a)
		if err != nil {
			t.Fatal(err)
		}

		if err := b.SetAttributes(f); err != nil {
			t.Fatal(err)
		}
	}

	grp, err := filepath.Glob(testDirAttributeGroup + wildcard)
	if err != nil {
		t.Fatal(err)
	}

	for _, g := range grp {
		f, err := os.Open(g)
		if err != nil {
			t.Fatal(err)
		}

		if err := b.SetAttributeGroups(f); err != nil {
			t.Fatal(err)
		}
	}

	i, err := b.Item()
	if err != nil {
		t.Error(i)
	}

	spew.Dump(i)
}

func TestMundaneBuilder_LinkRecipes(t *testing.T) {
	b, err := setupMundaneBuilder()
	if err != nil {
		t.Fatal(err)
	}

	if err := b.linkRecipes(); err != nil {
		t.Error(err)
	}

	r := b.Recipes["figurine"]
	for _, c := range r.Comps {
		for _, prop := range c.Properties {
			if len(prop.Attributes) != len(prop.AttributeNames) {
				t.Errorf("got: <%v>, want: <%v>", prop.Attributes, prop.AttributeNames)
			}

			if len(prop.AttributeGroups) != len(prop.AttributeGroupNames) {
				t.Errorf("got: <%v>, want: <%v>", prop.AttributeGroups, prop.AttributeGroupNames)
			}
		}
	}
}
