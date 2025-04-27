package entity

import "gorm.io/gorm"

// GORM hook: Validate before saving or creating
func (c *Collaboration) BeforeCreate(tx *gorm.DB) (err error) {
	if !c.Permission.IsValid() {
		return ErrInvalidPermission
	}
	return nil
}

func (c *Collaboration) BeforeSave(tx *gorm.DB) (err error) {
	if !c.Permission.IsValid() {
		return ErrInvalidPermission
	}
	return nil
}
