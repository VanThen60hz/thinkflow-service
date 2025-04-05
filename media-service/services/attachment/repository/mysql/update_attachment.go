package mysql

import (
	"context"

	"thinkflow-service/services/attachment/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateAttachment(ctx context.Context, id int, data *entity.Attachment) error {
	if err := repo.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
