package scene

import "press/schema"

type Entity struct {
	ID   string
	Name string
	Slug string

	Schema *schema.Entity
	Data   interface{}
}
