package models

import "github.com/google/uuid"

type PermissionMenu struct {
	PermissionID uuid.UUID  `gorm:"type:uuid;not null"`
	MenuID       uuid.UUID  `gorm:"type:uuid;not null"`
	Permission   Permission `gorm:"foreignKey:PermissionID;references:ID"`
	Menu         Menu       `gorm:"foreignKey:MenuID;references:ID"`
}

func (PermissionMenu) TableName() string {
	return "permissions_menus"
}
