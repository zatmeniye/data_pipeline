package mapper

import (
	"bi/internal/dto"
	"bi/internal/entity"
	"context"
)

type SourceTypMapper struct {
}

func NewSourceTypMapper() *SourceTypMapper {
	return &SourceTypMapper{}
}

func (m *SourceTypMapper) EntityToDto(
	ctx context.Context,
	typ entity.SourceTyp,
) (dto.SourceTypDto, error) {
	return dto.SourceTypDto{
		Id:   typ.Id,
		Name: typ.Name,
	}, nil
}
