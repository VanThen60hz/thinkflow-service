package mysql

import (
	"context"

	"thinkflow-service/services/text/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteText(ctx context.Context, id int) error {
	if err := repo.db.Table(entity.Text{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
