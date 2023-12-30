package service

import (
	"bi/internal/entity"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type SourceTypService struct {
	l                   *zap.SugaredLogger
	sourceTypRepository SourceTypRepository
}

func NewSourceTypService(
	l *zap.SugaredLogger,
	sourceTypRepository SourceTypRepository,
) *SourceTypService {
	return &SourceTypService{sourceTypRepository: sourceTypRepository, l: l}
}

func (s *SourceTypService) GetAll(ctx context.Context) ([]entity.SourceTyp, error) {
	sourceTypes, err := s.sourceTypRepository.GetAll(ctx)
	if err != nil {
		s.l.Error(err)
		return nil, fiber.NewError(
			http.StatusInternalServerError,
			"произошла ошибка при получении типов источников",
		)
	}
	return sourceTypes, nil
}
