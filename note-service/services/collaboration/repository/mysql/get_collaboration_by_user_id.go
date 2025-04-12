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
		var total int64
		if err := db.Count(&total).Error; err != nil {
			return nil, errors.WithStack(err)
		}
		paging.Total = int64(total)

		db = db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit)
	}

	if err := db.Find(&collaborations).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	for i := range collaborations {
		collaborations[i].Mask()
	}

	return collaborations, nil
}
