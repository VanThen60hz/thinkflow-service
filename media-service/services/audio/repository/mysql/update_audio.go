package mysql

import (
	"context"
	"fmt"

	"thinkflow-service/services/audio/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) UpdateAudio(ctx context.Context, id int, data *entity.AudioDataUpdate) error {
	fmt.Println("data", data)
	if err := repo.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
