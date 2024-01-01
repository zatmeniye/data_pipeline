package repository

import (
	"bi/internal/entity"
	"bi/pkg/database"
	"context"
	"database/sql"
)

type SourceRepository struct {
	db *database.Database
}

func NewSourceRepository(db *database.Database) *SourceRepository {
	return &SourceRepository{db: db}
}

func (r *SourceRepository) GetAll(ctx context.Context) ([]entity.Source, error) {
	query, args, err := r.db.B.
		Select("source_id", "dsn", "source_typ_id", "name").
		From("source").
		Join("source_typ USING (source_typ_id)").
		OrderBy("source_id").
		ToSql()
	if err != nil {
		return make([]entity.Source, 0, 0), err
	}

	var rows *sql.Rows

	if rows, err = r.db.Query(ctx, query, args...); err != nil {
		return make([]entity.Source, 0, 0), err
	}

	defer func() { _ = rows.Close() }()

	sources := make([]entity.Source, 0)

	for rows.Next() {
		var source entity.Source

		if err = rows.Scan(&source.Id, &source.Dsn, &source.Typ.Id, &source.Typ.Name); err != nil {
			return make([]entity.Source, 0, 0), err
		}

		sources = append(sources, source)
	}

	return sources, nil
}

func (r *SourceRepository) Add(ctx context.Context, source entity.Source) (uint32, error) {
	query, args, err := r.db.B.
		Insert("source").
		Columns("dsn", "source_typ_id").
		Values(source.Dsn, source.Typ.Id).
		Suffix("RETURNING source_id").
		ToSql()
	if err != nil {
		return 0, err
	}

	var sourceId uint32

	return sourceId, r.db.QueryRow(ctx, query, args...).Scan(&sourceId)
}
