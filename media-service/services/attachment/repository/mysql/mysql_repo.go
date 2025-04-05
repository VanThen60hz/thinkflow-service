package mysql

import (
	"context"
	"errors"

	"thinkflow-service/services/attachment/entity"

	"github.com/VanThen60hz/service-context/core"
	"gorm.io/gorm"
)

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{db: db}
}

func (repo *mysqlRepo) GetByID(ctx context.Context, id int64) (*entity.Attachment, error) {
	var attachment entity.Attachment
	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&attachment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrNotFound
		}
		return nil, err
	}
	return &attachment, nil
}

func (repo *mysqlRepo) GetByNoteID(ctx context.Context, noteID int64) ([]entity.Attachment, error) {
	var attachments []entity.Attachment
	if err := repo.db.WithContext(ctx).Where("note_id = ?", noteID).Find(&attachments).Error; err != nil {
		return nil, err
	}
	return attachments, nil
}
