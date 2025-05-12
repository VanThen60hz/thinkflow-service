package repository

import (
	"context"

	"thinkflow-service/services/notification/model"

	"gorm.io/gorm"
)

type FCMTokenRepository struct {
	db *gorm.DB
}

func NewFCMTokenRepository(db *gorm.DB) *FCMTokenRepository {
	return &FCMTokenRepository{db: db}
}

func (r *FCMTokenRepository) Save(ctx context.Context, token *model.FCMToken) error {
	return r.db.WithContext(ctx).Save(token).Error
}

func (r *FCMTokenRepository) GetByUserID(ctx context.Context, userID string) ([]model.FCMToken, error) {
	var tokens []model.FCMToken
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&tokens).Error
	return tokens, err
}

func (r *FCMTokenRepository) Delete(ctx context.Context, token string) error {
	return r.db.WithContext(ctx).Where("token = ?", token).Delete(&model.FCMToken{}).Error
}
