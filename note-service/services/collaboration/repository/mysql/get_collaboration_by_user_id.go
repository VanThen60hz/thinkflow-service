package mysql

import (
	"context"

	"thinkflow-service/services/collaboration/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) GetCollaborationByUserId(ctx context.Context, userId int) ([]entity.Collaboration, error) {
	var collaborations []entity.Collaboration

	db := repo.db.Table(entity.Collaboration{}.TableName())

	if err := db.Where("user_id = ?", userId).Find(&collaborations).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	for i := range collaborations {
		collaborations[i].Mask()
	}

	return collaborations, nil
}
