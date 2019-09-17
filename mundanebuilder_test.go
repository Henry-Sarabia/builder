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
	testDirRecipe = "testdata/recipes"
	testDirAttribute = "testdata/properties/types"
	testDirAttributeGroup = "testdata/properties/groups"
)

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