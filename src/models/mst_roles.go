package models

import (
	"auth-rbac/src/utils/enums"
	"time"

	"github.com/google/uuid"
)

func (Role) TableName() string {
	return "mst_role"
}

type Role struct {
	ID        uuid.UUID          `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string             `gorm:"size:20;not null"`
	Status    enums.StatusActive `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
