package mysql

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *mysqlRepo) GetUsersByIds(ctx context.Context, ids []int) ([]entity.User, error) {
	var result []entity.User

	if err := repo.db.
		Table(entity.User{}.TableName()).
		Where("id in (?)", ids).
		Find(&result).Error; err != nil {
		return nil, errors.Wrap(err, entity.ErrCannotGetUser.Error())
	}

	return result, nil
}

func (repo *mysqlRepo) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	var data entity.User

	if err := repo.db.
		Table(data.TableName()).
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, core.ErrNotFound
		}
		return nil, errors.Wrap(err, entity.ErrCannotGetUser.Error())
	}

	return &data, nil
}
