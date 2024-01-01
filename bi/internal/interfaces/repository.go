package interfaces

import (
	"bi/internal/entity"
	"context"
)

type SourceTypRepository interface {
	GetAll(context.Context) ([]entity.SourceTyp, error)
	GetById(context.Context, uint32) (entity.SourceTyp, error)
}

type SourceRepository interface {
	GetAll(context.Context) ([]entity.Source, error)
	Add(context.Context, entity.Source) (uint32, error)
}
