package mysql

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdatePassword(ctx context.Context, email, salt, hashedPassword string) error {
	if err := repo.db.Table(entity.Auth{}.TableName()).
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"salt":     salt,
			"password": hashedPassword,
		}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}