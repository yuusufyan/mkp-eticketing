package models

import "github.com/google/uuid"

type RolesPermissions struct {
	RoleID       uuid.UUID  `gorm:"type:uuid;not null"`
	PermissionID uuid.UUID  `gorm:"type:uuid;not null"`
	Role         Role       `gorm:"foreignKey:RoleID;references:ID"`
	Permission   Permission `gorm:"foreignKey:PermissionID;references:ID"`
}

func (RolesPermissions) TableName() string {
	return "roles_permissions"
}
