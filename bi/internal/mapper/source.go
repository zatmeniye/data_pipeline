package mapper

import (
	"bi/internal/dto"
	"bi/internal/entity"
	"bi/internal/interfaces"
	"context"
)

type SourceMapper struct {
	sourceTypMapper     interfaces.SourceTypMapper
	sourceTypRepository interfaces.SourceTypRepository
}

func NewSourceMapper(
	sourceTypMapper interfaces.SourceTypMapper,
	sourceTypRepository interfaces.SourceTypRepository,
) *SourceMapper {
	return &SourceMapper{
		sourceTypMapper:     sourceTypMapper,
		sourceTypRepository: sourceTypRepository,
	}
}

func (m *SourceMapper) EntityToDto(
	ctx context.Context,
	source entity.Source,
) (dto.SourceDto, error) {
	typ, err := m.sourceTypMapper.EntityToDto(ctx, source.Typ)
	if err != nil {
		return dto.SourceDto{}, err
	}
	return dto.SourceDto{
		Id:  source.Id,
		Dsn: source.Dsn,
		Typ: typ,
	}, nil
}

func (m *SourceMapper) AddDtoToEntity(
	ctx context.Context,
	addDto dto.SourceAddDto,
) (entity.Source, error) {
	typ, err := m.sourceTypRepository.GetById(ctx, addDto.TypId)
	if err != nil {
		return entity.Source{}, err
	}
	return entity.Source{
		Dsn: addDto.Dsn,
		Typ: typ,
	}, nil
}
