package mysql

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *mysqlRepo) GetAuth(ctx context.Context, email string) (*entity.Auth, error) {
	var data entity.Auth

	if err := repo.db.
		Table(data.TableName()).
		Where("email = ?", email).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrRecordNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &data, nil
}
