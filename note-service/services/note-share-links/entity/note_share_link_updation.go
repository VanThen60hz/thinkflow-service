package entity

import "time"

type NoteShareLinkUpdate struct {
	ExpiresAt  *time.Time `json:"expires_at,omitempty" gorm:"column:expires_at"`
	Permission *string    `json:"permission,omitempty" gorm:"column:permission;type:enum('read','write')"`
}

func (NoteShareLinkUpdate) TableName() string { return NoteShareLink{}.TableName() }

func (link *NoteShareLinkUpdate) Validate() error {
	if link.Permission != nil {
		if *link.Permission != "read" && *link.Permission != "write" {
			return ErrInvalidPermission
		}
	}
	return nil
}
