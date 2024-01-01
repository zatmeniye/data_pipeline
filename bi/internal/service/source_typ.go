package service

import (
	"bi/internal/dto"
	"bi/internal/interfaces"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type SourceTypService struct {
	l                   *zap.SugaredLogger
	sourceTypRepository interfaces.SourceTypRepository
	sourceTypMapper     interfaces.SourceTypMapper
}

func NewSourceTypService(
	l *zap.SugaredLogger,
	sourceTypRepository interfaces.SourceTypRepository,
	sourceTypMapper interfaces.SourceTypMapper,
) *SourceTypService {
	return &SourceTypService{
		l:                   l,
		sourceTypRepository: sourceTypRepository,
		sourceTypMapper:     sourceTypMapper,
	}
}

func (s *SourceTypService) GetAll(ctx context.Context) ([]dto.SourceTypDto, error) {
	sourceTypes, err := s.sourceTypRepository.GetAll(ctx)
	if err != nil {
		s.l.Error(err)
		return nil, fiber.NewError(
			http.StatusInternalServerError,
			"произошла ошибка при получении типов источников",
		)
	}
	dtos := make([]dto.SourceTypDto, 0, len(sourceTypes))
	for _, sourceType := range sourceTypes {
		sourceTypeDto, err := s.sourceTypMapper.EntityToDto(ctx, sourceType)
		if err != nil {
			s.l.Error(err)
			return nil, fiber.NewError(
				http.StatusInternalServerError,
				"произошла ошибка при получении типов источников",
			)
		}
		dtos = append(dtos, sourceTypeDto)
	}
	return dtos, nil
}
