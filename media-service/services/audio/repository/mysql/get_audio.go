package mysql

import (
	"context"

	"thinkflow-service/services/audio/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *mysqlRepo) GetAudioById(ctx context.Context, id int) (*entity.Audio, error) {
	var data entity.Audio

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

func (repo *mysqlRepo) ListAudios(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Audio, error) {
	var result []entity.Audio

	db := repo.db.Table(entity.Audio{}.TableName())

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

func (repo *mysqlRepo) GetAudiosByNoteId(ctx context.Context, noteId int) ([]entity.Audio, error) {
	var result []entity.Audio

	if err := repo.db.
		Table(entity.Audio{}.TableName()).
		Where("note_id = ?", noteId).
		Find(&result).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
