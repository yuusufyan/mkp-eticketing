package models

import (
	"github.com/google/uuid"
)

func (UsersRoles) TableName() string {
	return "users_roles"
}

type UsersRoles struct {
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	RoleID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID;references:ID"`
	Role   Role      `gorm:"foreignKey:RoleID;references:ID"`
}
