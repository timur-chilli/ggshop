package PGStorage

import (
	"context"
	// "log"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (storage *PGStorage) InsertGGOrderInfo(ctx context.Context, originalInfos []*models.GGOrderInfo) error {
	query := storage.insertQuery(originalInfos)
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

func (storage *PGStorage) insertQuery(originalInfos []*models.GGOrderInfo) sq.Sqlizer {
	infos := lo.Map(originalInfos, func(info *models.GGOrderInfo, _ int) *GGOrderInfo {
		return &GGOrderInfo{
			CustomerName: info.CustomerName,
			Email:        info.Email,
			Details:      info.Details,
		}
	})

	q := sq.Insert(tableName).Columns(CustomerNameСolumnName, EmailСolumnName, DetailsСolumnName).
		PlaceholderFormat(sq.Dollar)
	for _, info := range infos {
		q = q.Values(info.CustomerName, info.Email, info.Details)
	}
	return q
}

// func (storage *PGStorage) UpdateGGOrderInfo(ctx context.Context, originalInfos []*models.GGOrderInfo) error {
// 	var err error = nil
// 	for _, info := range originalInfos {
// 		query := sq.Update(tableName).Where(sq.Eq{IDСolumnName: &info.ID})
// 		query.SetMap(map[string]interface{}{CustomerNameСolumnName: &info.CustomerName, EmailСolumnName: &info.Email, DetailsСolumnName: &info.Details})
// 		queryText, args, err := query.ToSql()
// 		if err != nil {
// 			return errors.Wrap(err, "ошибка при генерации SQL-запроса")
// 		}
// 		_, err = storage.db.Exec(ctx, queryText, args...)
// 		if err != nil {
// 			err = errors.Wrap(err, "ошибка при выполнении SQL-запроса")
// 		}
// 		log.Printf("postgres: succesfully updated %v with %v %v %v", info.ID, info.CustomerName, info.Email, info.Details)
// 	}
// 	return err
// }
