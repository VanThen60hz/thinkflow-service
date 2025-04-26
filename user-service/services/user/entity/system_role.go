package entity

type SystemRole string

const (
	RoleSuperAdmin SystemRole = "sadmin"
	RoleAdmin      SystemRole = "admin"
	RoleUser       SystemRole = "user"
)
