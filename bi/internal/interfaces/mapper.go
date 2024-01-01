package interfaces

import (
	"bi/internal/dto"
	"bi/internal/entity"
	"context"
)

type SourceMapper interface {
	EntityToDto(context.Context, entity.Source) (dto.SourceDto, error)
	AddDtoToEntity(context.Context, dto.SourceAddDto) (entity.Source, error)
}

type SourceTypMapper interface {
	EntityToDto(context.Context, entity.SourceTyp) (dto.SourceTypDto, error)
}
