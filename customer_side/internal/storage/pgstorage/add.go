package PGStorage

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (storage *PGStorage) InsertGGOrderInfo(ctx context.Context, GGOrderInfos []*models.GGOrderInfo) error {
	query := storage.insertQuery(GGOrderInfos)
	queryText, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "ошибка при генерации SQL-запроса")
	}
	_, err = storage.db.Exec(ctx, queryText, args...)
	if err != nil {
		err = errors.Wrap(err, "ошибка при выполнении SQL-запроса")
	}
	return err
}

func (storage *PGStorage) insertQuery(GGOrderInfos []*models.GGOrderInfo) squirrel.Sqlizer {
	infos := lo.Map(GGOrderInfos, func(info *models.GGOrderInfo, _ int) *GGOrderInfo {
		return &GGOrderInfo{
			CustomerName: info.CustomerName,
			Email:        info.Email,
			Details:      info.Details,
		}
	})

	q := squirrel.Insert(tableName).Columns(CustomerNameСolumnName, EmailСolumnName, DetailsСolumnName).
		PlaceholderFormat(squirrel.Dollar)
	for _, info := range infos {
		q = q.Values(info.CustomerName, info.Email, info.Details)
	}
	return q
}
