package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common columns for all tables
type Base struct {
	ID        string    `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index;default:NULL" json:"-" validate:"omitempty"`
	IsDeleted bool      `json:"isDeleted" default:"false"`
}

// BeforeCreate will set a UUID rather than numeric ID
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	if base.ID == "" {
		base.ID = uuid.New().String()
	}
	return nil
}

// Job status constants
type JobStatus string

const (
	JobStatusQueued     JobStatus = "QUEUED"
	JobStatusProcessing JobStatus = "PROCESSING"
	JobStatusCompleted  JobStatus = "COMPLETED"
	JobStatusFailed     JobStatus = "FAILED"
	JobStatusCancelled  JobStatus = "CANCELLED"
)

type UserRole string

const (
	UserRoleSuperAdmin UserRole = "SUPER_ADMIN"
	UserRoleAdmin      UserRole = "ADMIN"
	UserRoleMember     UserRole = "MEMBER"
)

type InviteStatus string

const (
	InviteStatusPending  InviteStatus = "PENDING"
	InviteStatusAccepted InviteStatus = "ACCEPTED"
	InviteStatusRejected InviteStatus = "REJECTED"
)
