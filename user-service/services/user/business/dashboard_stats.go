package business

import (
	"context"
	"time"

	"github.com/VanThen60hz/service-context/core"
)

type DashboardStats struct {
	TotalUsers    int64        `json:"total_users"`
	TotalNotes    int64        `json:"total_notes"`
	ActiveUsers   int64        `json:"active_users"`
	NewUsersToday int64        `json:"new_users_today"`
	SystemHealth  SystemHealth `json:"system_health"`
}

type SystemHealth struct {
	Status      string    `json:"status"`
	LastChecked time.Time `json:"last_checked"`
}

func (biz *business) GetDashboardStats(ctx context.Context) (*DashboardStats, error) {
	requester := core.GetRequester(ctx)
	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID())

	requesterDetail, err := biz.GetUserDetails(ctx, requesterId)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("cannot get requester status").
			WithDebug(err.Error())
	}

	requesterRole := requesterDetail.SystemRole
	if requesterRole != "admin" && requesterRole != "sadmin" {
		return nil, core.ErrForbidden.
			WithError("cannot access: user is not an admin or super admin").
			WithDebug("requester role: " + string(requesterRole))
	}

	totalUsers, err := biz.userRepo.CountUsers(ctx)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("cannot get total users").
			WithDebug(err.Error())
	}

	activeUsers, err := biz.userRepo.CountUsersByStatus(ctx, "active")
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("cannot get active users").
			WithDebug(err.Error())
	}

	today := time.Now().Truncate(24 * time.Hour)
	newUsersToday, err := biz.userRepo.CountUsersCreatedAfter(ctx, today)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("cannot get new users today").
			WithDebug(err.Error())
	}

	totalNotes, err := biz.noteRepo.CountNotes(ctx)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError("cannot get total notes").
			WithDebug(err.Error())
	}

	return &DashboardStats{
		TotalUsers:    totalUsers - 1,
		TotalNotes:    totalNotes,
		ActiveUsers:   activeUsers - 1,
		NewUsersToday: newUsersToday,
		SystemHealth: SystemHealth{
			Status:      "healthy",
			LastChecked: time.Now(),
		},
	}, nil
}
