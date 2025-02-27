package mysql

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
)

func (repo *mysqlRepo) ListNotes(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Note, error) {
	var notes []entity.Note

	db := repo.db.Table(entity.Note{}.TableName())

	if userId := filter.UserId; userId != nil {
		uid, err := core.FromBase58(*userId)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		db = db.Where("user_id = ?", uid.GetLocalID())
	}

	// Count total records match conditions
	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	// Query data with paging
	if err := db.Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&notes).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return notes, nil
}
