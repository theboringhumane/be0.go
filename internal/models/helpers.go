package models

import (
	"gorm.io/gorm"
)

// GetTeamByName retrieves a team from the database by its name
func GetTeamByName(name string, db *gorm.DB) (*Team, error) {
	team := &Team{}
	if err := db.Where("name = ? AND is_deleted = false", name).First(team).Error; err != nil {
		return nil, err
	}
	return team, nil
}

func GetFileByID(id string, db *gorm.DB) (*File, error) {
	file := &File{}
	if err := db.Where("id = ? AND is_deleted = false", id).First(file).Error; err != nil {
		return nil, err
	}
	return file, nil
}
