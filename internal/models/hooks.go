package models

import (
	"be0/internal/events"

	"gorm.io/gorm"
)

func (t *TeamInvite) AfterCreate(tx *gorm.DB) error {
	log.Info("Team invite created %v", t)
	events.Emit("invite.created", t)
	return nil
}
