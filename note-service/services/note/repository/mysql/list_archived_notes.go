package mysql

import (
	"context"

	"thinkflow-service/services/note/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
)

func (repo *mysqlRepo) ListArchivedNotes(ctx context.Context, filter *entity.Filter, paging *core.Paging) ([]entity.Note, error) {
	var notes []entity.Note

	db := repo.db.Table(entity.Note{}.TableName())

	if userId := filter.UserId; userId != nil {
		uid, err := core.FromBase58(*userId)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		db = db.Where("user_id = ?", uid.GetLocalID())
	}

	// Only show archived notes
	db = db.Where("archived = ?", true)

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
		Find(&notes).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	if len(notes) > 0 {
		notes[len(notes)-1].Mask()
		paging.NextCursor = notes[len(notes)-1].FakeId.String()
	}

	return notes, nil
}
