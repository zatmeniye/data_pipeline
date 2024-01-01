package repository

import (
	"bi/internal/entity"
	"bi/pkg/database"
	"context"
	"database/sql"
)

type SourceTypRepository struct {
	db *database.Database
}

func NewSourceTypRepository(db *database.Database) *SourceTypRepository {
	return &SourceTypRepository{db: db}
}

func (r *SourceTypRepository) GetAll(ctx context.Context) ([]entity.SourceTyp, error) {
	query, args, err := r.db.B.
		Select("source_typ_id", "name").
		From("source_typ").
		OrderBy("source_typ_id").
		ToSql()
	if err != nil {
		return make([]entity.SourceTyp, 0), err
	}

	var rows *sql.Rows

	if rows, err = r.db.Query(ctx, query, args...); err != nil {
		return make([]entity.SourceTyp, 0), err
	}
	defer func() { _ = rows.Close() }()

	sourceTypes := make([]entity.SourceTyp, 0)

	for rows.Next() {
		var sourceTyp entity.SourceTyp
		if err = rows.Scan(&sourceTyp.Id, &sourceTyp.Name); err != nil {
			return make([]entity.SourceTyp, 0), err
		}
		sourceTypes = append(sourceTypes, sourceTyp)
	}

	return sourceTypes, nil
}

func (r *SourceTypRepository) GetById(
	ctx context.Context,
	id uint32,
) (entity.SourceTyp, error) {
	query, args, err := r.db.B.
		Select("source_typ_id", "name").
		From("source_typ").
		Where("source_typ_id = ?", id).
		ToSql()
	if err != nil {
		return entity.SourceTyp{}, err
	}

	var sourceTyp entity.SourceTyp

	return sourceTyp, r.db.QueryRow(ctx, query, args...).Scan(
		&sourceTyp.Id,
		&sourceTyp.Name,
	)
}
