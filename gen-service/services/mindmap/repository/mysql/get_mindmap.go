package mysql

import (
	"context"

	"thinkflow-service/services/mindmap/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *mysqlRepo) GetMindmapById(ctx context.Context, id int) (*entity.Mindmap, error) {
	var data entity.Mindmap

	if err := repo.db.
		Table(data.TableName()).
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.ErrRecordNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &data, nil
}
