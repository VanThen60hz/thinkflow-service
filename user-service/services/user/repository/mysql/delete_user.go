package mysql

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteUser(ctx context.Context, id int) error {
	user := &entity.User{}
	
	if err := repo.db.Table(user.TableName()).Where("id = ?", id).Delete(user).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
} 