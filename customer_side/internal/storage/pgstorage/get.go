package PGStorage

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (storage *PGStorage) GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error) {
	query := storage.getQuery(IDs)
	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "generate query error")
	}
	rows, err := storage.db.Query(ctx, queryText, args...)
	if err != nil {
		return nil, errors.Wrap(err, "quering error")
	}
	var GGOrders []*models.GGOrderInfo
	for rows.Next() {
		var s models.GGOrderInfo
		if err := rows.Scan(&s.ID, &s.CustomerName, &s.Email, &s.Details); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		GGOrders = append(GGOrders, &s)
	}
	return GGOrders, nil
}

func (storage *PGStorage) getQuery(IDs []uint64) squirrel.Sqlizer {
	q := squirrel.Select(IDСolumnName, CustomerNameСolumnName, EmailСolumnName, DetailsСolumnName).From(tableName).
		Where(squirrel.Eq{IDСolumnName: IDs}).PlaceholderFormat(squirrel.Dollar)
	return q
}
