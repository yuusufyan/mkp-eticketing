package models

import (
	"github.com/google/uuid"
)

type RolePermissionMenu struct {
	ID             int        `gorm:"primaryKey"`
	RoleID         uuid.UUID  `gorm:"type:uuid;not null"`
	RoleName       string     `gorm:"not null"`
	PermissionName string     `gorm:"not null"`
	PermissionID   uuid.UUID  `gorm:"type:uuid;not null"`
	MenuID         uuid.UUID  `gorm:"type:uuid;not null"`
	MenuName       string     `gorm:"not null"`
	Role           Role       `gorm:"foreignKey:RoleID;references:ID"`
	Permission     Permission `gorm:"foreignKey:PermissionID;references:ID"`
	Menu           Menu       `gorm:"foreignKey:MenuID;references:ID"`
}

func (RolePermissionMenu) TableName() string {
	return "roles_permissions_menus"
}
