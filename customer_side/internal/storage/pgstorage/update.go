package PGStorage

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (storage *PGStorage) UpdateGGOrderInfo(ctx context.Context, originalInfos []*models.GGOrderInfo) error {
	query := storage.updateQuery(originalInfos)
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

func (storage *PGStorage) updateQuery(originalInfos []*models.GGOrderInfo) squirrel.Sqlizer {
	infos := lo.Map(originalInfos, func(info *models.GGOrderInfo, _ int) *GGOrderInfo {
		return &GGOrderInfo{
			ID: info.ID,
			CustomerName: info.CustomerName,
			Email:        info.Email,
			Details:      info.Details,
		}
	})
	
	q := squirrel.Update(tableName).
		PlaceholderFormat(squirrel.Dollar)
	for _, info := range infos {
		tempMap := map[string]any{CustomerNameСolumnName: info.CustomerName, EmailСolumnName: info.Email, DetailsСolumnName: info.Details}
		q = q.SetMap(tempMap).Where(squirrel.Eq{"ID": info.ID})
	}
	return q
}
