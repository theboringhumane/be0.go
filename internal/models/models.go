package models

import (
	"be0/internal/events"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	Base
	Name    string       `gorm:"not null" json:"name" validate:"required,min=2"`
	Users   []User       `gorm:"foreignKey:TeamID;references:ID" json:"users,omitempty"`
	Invites []TeamInvite `gorm:"foreignKey:TeamID;references:ID;constraint:OnDelete:CASCADE" json:"invites,omitempty"`
}

func (t *Team) BeforeCreate(tx *gorm.DB) error {
	if t.ID == "" {
		t.ID = uuid.New().String()
	}
	return nil
}

func (t *Team) AfterCreate(tx *gorm.DB) error {
	// Emit team created event
	events.Emit("team.created", t)
	return nil
}

type TeamInvite struct {
	Base
	Email     string       `gorm:"not null" json:"email" validate:"required,email"`
	Name      string       `gorm:"not null" json:"name" validate:"required,min=2"`
	TeamID    string       `gorm:"type:uuid;not null" json:"teamId" validate:"required,uuid"`
	Team      *Team        `json:"team,omitempty"`
	InviterID string       `gorm:"type:uuid;not null" json:"inviterId" validate:"required,uuid"`
	Inviter   *User        `json:"inviter,omitempty"`
	Role      UserRole     `gorm:"not null;default:'MEMBER'" json:"role" validate:"required,oneof=MEMBER ADMIN"`
	Code      string       `gorm:"not null" json:"code" validate:"required=min=4"`
	Status    InviteStatus `gorm:"not null;default:'PENDING'" json:"status" validate:"required,oneof=PENDING ACCEPTED REJECTED"`
	ExpiresAt time.Time    `gorm:"not null" json:"expiresAt" validate:"required,gt=now"`
}

type File struct {
	Base
	TeamID    string `gorm:"type:uuid" json:"teamId" validate:"omitempty,uuid"`
	Team      *Team  `json:"team,omitempty"`
	Path      string `gorm:"not null" json:"path" validate:"required"`
	UserID    string `gorm:"type:uuid;default:NULL" json:"userId" validate:"omitempty,uuid"`
	User      *User  `json:"user,omitempty"`
	Name      string `gorm:"not null" json:"name" validate:"required"`
	Size      int64  `gorm:"not null" json:"size" validate:"required,min=1"`
	Type      string `gorm:"not null" json:"type" validate:"required"`
	SignedURL string `gorm:"-" json:"signedUrl,omitempty"` // Virtual field
}

func (f *File) BeforeCreate(tx *gorm.DB) error {
	if f.ID == "" {
		f.ID = uuid.New().String()
	}
	return nil
}

func (f *File) AfterFind(tx *gorm.DB) error {
	registryMu.RLock()
	generator := urlGenerator
	registryMu.RUnlock()

	if generator != nil {
		// Generate URL with 1-hour expiry
		url, err := generator.GetSignedURL(tx.Statement.Context, f.Path, time.Hour)
		if err != nil {
			return fmt.Errorf("failed to generate signed URL: %w", err)
		}
		f.SignedURL = url
	}
	return nil
}

// IsValidUserRole checks if a given role is valid
func IsValidUserRole(role UserRole) bool {
	switch role {
	case UserRoleAdmin, UserRoleMember, UserRoleSuperAdmin:
		return true
	default:
		return false
	}
}
