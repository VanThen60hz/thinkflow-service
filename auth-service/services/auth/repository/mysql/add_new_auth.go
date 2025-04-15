package mysql

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) AddNewAuth(ctx context.Context, data *entity.Auth) error {
	if err := repo.db.Table("abctest").Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
