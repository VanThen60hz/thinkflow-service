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
	var audios []entity.Audio

	db := repo.db.Table(entity.Audio{}.TableName())

	if noteId := filter.NoteID; noteId != nil {
		db = db.Where("note_id = ?", *noteId)
	}

	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := core.FromBase58(v)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Select("*").
		Limit(paging.Limit).
		Order("id desc").
		Find(&audios).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	if len(audios) > 0 {
		audios[len(audios)-1].Mask()
		paging.NextCursor = audios[len(audios)-1].FakeId.String()
	}

	return audios, nil
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
