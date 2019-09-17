package builder

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

type AttributeGroup struct {
	Name           string   `json:"name"`
	AttributeNames []string `json:"attribute_names"`
	Attributes     []*Attribute
}

// readAttributeGroup reads the JSON-encoded AttributeGroups from the provided
// Reader.
func readAttributeGroup (r io.Reader) ([]AttributeGroup, error) {
	var attr []AttributeGroup

	if err := json.NewDecoder(r).Decode(&attr); err != nil {
		return nil, errors.Wrap(err, "cannot decode AttributeGroup from io.Reader")
	}

	return attr, nil
}