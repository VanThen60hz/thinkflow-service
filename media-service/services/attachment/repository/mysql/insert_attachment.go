package mysql

import (
	"context"

	"thinkflow-service/services/attachment/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) AddNewAttachment(ctx context.Context, data *entity.AttachmentCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
