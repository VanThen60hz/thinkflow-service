package mysql

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) DeleteAuth(ctx context.Context, email string) error {
	auth := &entity.Auth{}
	
	if err := repo.db.Table(auth.TableName()).Where("email = ?", email).Delete(auth).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
} 