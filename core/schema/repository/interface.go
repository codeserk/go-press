package repository

import (
	"context"
	"press/schema"
)

type Interface interface {
	Create(context context.Context, schema schema.Entity) error
	Find(context context.Context) ([]*schema.Entity, error)
}
