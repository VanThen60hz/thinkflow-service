package entity

type Status string

const (
	StatusActive        Status = "active"
	StatusPendingVerify Status = "waiting_verify"
	StatusBanned        Status = "banned"
)
