package keeper

import (
	"bi/config"
	"bi/pkg/database"
	"context"
	"database/sql"
)

// Keeper - ОБОЛОЧКА БАЗЫ ДАННЫХ,
// КОТОРАЯ ПОЗВОЛЯЕТ ВЫПОЛНЯТЬ ПРОИЗВОЛЬНЫЕ ЗАПРОСЫ
// И ПОЛУЧАТЬ НАБОРЫ ДАННЫХ.
type Keeper struct {
	db *database.Database
}

func New(cfg config.Database) (*Keeper, error) {
	db, err := database.New(cfg)
	if err != nil {
		return nil, err
	}
	return &Keeper{db: db}, nil
}

func (k *Keeper) Close() {
	k.db.Close()
}

// Get - ПРИНИМАЕТ SQL ЗАПРОС,
// ВЫПОЛНЯЕТ ЕГО И ВОЗВРАЩАЕТ НАБОР ДАННЫХ.
func (k *Keeper) Get(
	ctx context.Context,
	query string,
) (Dataset, error) {

	rows, err := k.db.Query(ctx, query, make([]any, 0, 0)...)
	if err != nil {
		return Dataset{}, err
	}
	defer func() { _ = rows.Close() }()

	columnTypes := make([]*sql.ColumnType, 0)

	if columnTypes, err = rows.ColumnTypes(); err != nil {
		return Dataset{}, err
	}

	var dataset Dataset

	for _, columnType := range columnTypes {
		dataset.Columns = append(dataset.Columns, Column{
			Typ:  columnType.ScanType().String(),
			Name: columnType.Name(),
		})
	}

	for rows.Next() {
		pointers := make([]any, len(dataset.Columns), len(dataset.Columns))
		for i := range dataset.Columns {
			pointers[i] = new(*any)
		}

		if err = rows.Scan(pointers...); err != nil {
			return Dataset{}, err
		}

		row := make(map[string]any, len(dataset.Columns))
		for i := range dataset.Columns {
			row[dataset.Columns[i].Name] = pointers[i]
		}
		dataset.Rows = append(dataset.Rows, row)
	}

	return dataset, nil
}
