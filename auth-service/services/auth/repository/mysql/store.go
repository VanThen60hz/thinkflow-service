package mysql

import (
	"context"

	"thinkflow-service/services/auth/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{db: db}
}

func (repo *mysqlRepo) AddNewAuth(ctx context.Context, data *entity.Auth) error {
	if err := repo.db.Table(data.TableName()).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

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
