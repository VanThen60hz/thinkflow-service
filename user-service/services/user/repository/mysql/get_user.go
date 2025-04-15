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

func (repo *mysqlRepo) GetUserIdByEmail(ctx context.Context, email string) (int, error) {
	var userId int

	if err := repo.db.
		Table(entity.User{}.TableName()).
		Select("id").
		Where("email = ?", email).
		First(&userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, core.ErrNotFound
		}
		return 0, errors.Wrap(err, entity.ErrCannotGetUser.Error())
	}

	return userId, nil
}
