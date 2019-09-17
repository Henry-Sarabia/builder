package builder

import "io"

type Builder interface {
	SetRecipes(io.Reader) error
	SetAttributes(io.Reader) error
	SetAttributeGroups(io.Reader) error
	LinkResources() error
	Item() (*Item, error)
}

type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Class       string  `json:"class"`
	Value       float64 `json:"value"`
	Weight      float64 `json:"weight"`
	Description string  `json:"description"`
}