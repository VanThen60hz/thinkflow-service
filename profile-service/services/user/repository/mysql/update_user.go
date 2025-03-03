package mysql

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateUser(ctx context.Context, id int, data *entity.UserDataUpdate) error {
	if err := repo.db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
