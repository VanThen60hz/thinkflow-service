package mysql

import (
	"context"

	"thinkflow-service/services/mindmap/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateMindmap(ctx context.Context, id int, data *entity.MindmapDataUpdate) error {
	var mindmap entity.Mindmap
	if err := repo.db.First(&mindmap, id).Error; err != nil {
		return errors.WithStack(err)
	}

	mindmap.MindmapData = data.MindmapData
	if err := repo.db.Save(&mindmap).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
