package service

import (
	"bi/config"
	"bi/internal/dto"
	"bi/internal/entity"
	"bi/internal/interfaces"
	"bi/pkg/keeper"
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

type SourceService struct {
	l                *zap.SugaredLogger
	sourceRepository interfaces.SourceRepository
	sourceMapper     interfaces.SourceMapper
	keepers          map[uint32]*keeper.Keeper
}

func NewSourceService(
	l *zap.SugaredLogger,
	sourceRepository interfaces.SourceRepository,
	sourceMapper interfaces.SourceMapper,
) *SourceService {
	return &SourceService{
		l:                l,
		sourceRepository: sourceRepository,
		sourceMapper:     sourceMapper,
		keepers:          make(map[uint32]*keeper.Keeper),
	}
}

func (s *SourceService) Add(ctx context.Context, addDto dto.SourceAddDto) (uint32, error) {
	source, err := s.sourceMapper.AddDtoToEntity(ctx, addDto)
	if err != nil {
		s.l.Error(err)
		return 0, fiber.NewError(
			http.StatusInternalServerError,
			"произошла ошибка при добавлении источника",
		)
	}

	//TODO: СОЗДАТЬ ТРАНЗАКЦИЮ.
	source.Id, err = s.sourceRepository.Add(ctx, source)
	if err != nil {
		s.l.Error(err)
		return 0, fiber.NewError(
			http.StatusInternalServerError,
			"произошла ошибка при добавлении источника",
		)
	}

	if err = s.addKeeper(source); err != nil {
		//TODO: ОТКАТИТЬ ТРАНЗАКЦИЮ.
		s.l.Error(err)
		return 0, fiber.NewError(
			http.StatusInternalServerError,
			"произошла ошибка при добавлении источника",
		)
	}

	return source.Id, nil
}

func (s *SourceService) ExecuteQuery(
	ctx context.Context,
	query dto.QueryDto,
) (keeper.Dataset, error) {
	source, err := s.getKeeper(query.SourceId)
	if err != nil {
		return keeper.Dataset{}, fiber.NewError(
			http.StatusNotFound,
			err.Error(),
		)
	}

	var dataset keeper.Dataset

	if dataset, err = source.Get(ctx, query.Query); err != nil {
		s.l.Error(err)
		return keeper.Dataset{}, fiber.NewError(
			http.StatusInternalServerError,
			"произошла ошибка при выполнении запроса: "+err.Error(),
		)
	}

	return dataset, nil
}

func (s *SourceService) addKeeper(source entity.Source) error {
	k, err := keeper.New(config.Database{
		Type: source.Typ.Name,
		Dsn:  source.Dsn,
	})
	if err != nil {
		return err
	}

	s.keepers[source.Id] = k

	return nil
}

func (s *SourceService) deleteKeeper(sourceId uint32) {
	k, ok := s.keepers[sourceId]
	if !ok {
		k.Close()
		delete(s.keepers, sourceId)
	}
}

func (s *SourceService) getKeeper(sourceId uint32) (*keeper.Keeper, error) {
	k, ok := s.keepers[sourceId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("источника с идентификатором %d нет", sourceId))
	}
	return k, nil
}
