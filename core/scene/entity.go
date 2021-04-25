package scene

import "press/core/schema"

type Entity struct {
	ID   string
	Name string
	Slug string

	Schema *schema.Entity
	Data   interface{}
}
