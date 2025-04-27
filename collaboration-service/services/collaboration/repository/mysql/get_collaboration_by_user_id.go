package mysql

import (
	"context"

	"thinkflow-service/services/collaboration/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
)

func (repo *mysqlRepo) GetCollaborationByUserId(ctx context.Context, userId int, paging *core.Paging) ([]entity.Collaboration, error) {
	var collaborations []entity.Collaboration

	db := repo.db.Table(entity.Collaboration{}.TableName()).
		Where("user_id = ?", userId)

	if paging != nil {
		if err := db.Count(&paging.Total).Error; err != nil {
			return nil, errors.WithStack(err)
		}

		if v := paging.FakeCursor; v != "" {
			uid, err := core.FromBase58(v)
			if err != nil {
				return nil, errors.WithStack(err)
			}

			noteId := uid.GetLocalID()

			db = db.Where("note_id < ?", noteId)
		} else {
			db = db.Offset((paging.Page - 1) * paging.Limit)
		}

		db = db.Limit(paging.Limit)
	}

	if err := db.Order("note_id desc").Find(&collaborations).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return collaborations, nil
}
