package models

import (
	"auth-rbac/src/utils/enums"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID        uuid.UUID          `gorm:"type:uuid;primaryKey"`
	Name      string             `gorm:"size:20;not null"`
	Status    enums.StatusActive `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time          `gorm:"type:timestamp;not null"`
	CreatedBy string             `gorm:"size:20;not null"`
	UpdatedAt time.Time          `gorm:"type:timestamp;not null"`
	UpdatedBy string             `gorm:"size:20;not null"`
}

func (Role) TableName() string {
	return "mst_roles"
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return
}
