package mysql

import (
	"context"

	"thinkflow-service/services/text/entity"

	"github.com/pkg/errors"
)

func (repo *mysqlRepo) AddNewText(ctx context.Context, data *entity.TextDataCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
