package mysql

import (
	"context"

	"thinkflow-service/services/mindmap/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) AddNewMindmap(ctx context.Context, data *entity.MindmapDataCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
