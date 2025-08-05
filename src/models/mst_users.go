package models

import (
	"auth-rbac/src/utils/enums"
	"time"

	"github.com/google/uuid"
)

func (User) TableName() string {
	return "mst_user"
}

type User struct {
	ID        uuid.UUID          `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username  string             `gorm:"size:20;not null"`
	Email     string             `gorm:"size:100;not null"`
	Password  string             `gorm:"size:100;not null"`
	Status    enums.StatusActive `gorm:"type:varchar(20);not null"`
	RoleID    uuid.UUID
	Role      Role `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
