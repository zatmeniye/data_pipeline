package service

import (
	"bi/internal/entity"
	"context"
)

type SourceTypRepository interface {
	GetAll(context.Context) ([]entity.SourceTyp, error)
}
