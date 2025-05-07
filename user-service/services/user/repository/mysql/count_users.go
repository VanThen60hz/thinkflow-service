package mysql

import (
	"context"
	"time"

	"thinkflow-service/services/user/entity"
)

func (repo *mysqlRepo) CountUsers(ctx context.Context) (int64, error) {
	var count int64
	if err := repo.db.Model(&entity.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *mysqlRepo) CountUsersByStatus(ctx context.Context, status string) (int64, error) {
	var count int64
	if err := repo.db.Model(&entity.User{}).Where("status = ?", status).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *mysqlRepo) CountUsersCreatedAfter(ctx context.Context, time time.Time) (int64, error) {
	var count int64
	if err := repo.db.Model(&entity.User{}).Where("created_at >= ?", time).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
