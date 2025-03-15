package mysql

import (
	"context"

	"thinkflow-service/services/image/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *mysqlRepo) GetImageById(ctx context.Context, id int) (*entity.Image, error) {
	var data entity.Image

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

func (repo *mysqlRepo) ListImages(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Image, error) {
	var result []entity.Image

	db := repo.db.Table(entity.Image{}.TableName())

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
