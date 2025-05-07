package mysql

import (
	"context"

	"thinkflow-service/services/user/entity"

	"github.com/VanThen60hz/service-context/core"
	"github.com/pkg/errors"
)

func (repo *mysqlRepo) ListUsers(ctx context.Context, filter *entity.UserFilter, paging *core.Paging) ([]entity.User, error) {
	var users []entity.User
	db := repo.db.Table(entity.User{}.TableName())

	db = db.Where("system_role NOT IN ?", []string{"admin", "sadmin"})

	if email := filter.Email; email != nil && *email != "" {
		db = db.Where("email LIKE ?", "%"+*email+"%")
	}

	if firstName := filter.FirstName; firstName != nil && *firstName != "" {
		db = db.Where("first_name LIKE ?", "%"+*firstName+"%")
	}

	if lastName := filter.LastName; lastName != nil && *lastName != "" {
		db = db.Where("last_name LIKE ?", "%"+*lastName+"%")
	}

	if role := filter.Role; role != nil && *role != "" {
		db = db.Where("system_role = ?", *role)
	}

	if status := filter.Status; status != nil && *status != "" {
		db = db.Where("status = ?", *status)
	}

	if searchQuery := filter.SearchQuery; searchQuery != nil && *searchQuery != "" {
		query := "%" + *searchQuery + "%"
		db = db.Where("email LIKE ? OR first_name LIKE ? OR last_name LIKE ?", query, query, query)
	}

	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := core.FromBase58(v)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Select("*").
		Limit(paging.Limit).
		Order("id desc").
		Find(&users).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	if len(users) > 0 {
		users[len(users)-1].Mask()
		paging.NextCursor = users[len(users)-1].FakeId.String()
	}

	return users, nil
}
