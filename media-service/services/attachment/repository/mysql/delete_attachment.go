package mysql

import (
	"context"

	"thinkflow-service/services/attachment/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteAttachment(ctx context.Context, id int) error {
	if err := repo.db.Table(entity.Attachment{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
