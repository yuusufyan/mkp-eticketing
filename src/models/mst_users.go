package models

import (
	"auth-rbac/src/utils/enums"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (User) TableName() string {
	return "mst_users"
}

type User struct {
	ID        uuid.UUID          `gorm:"type:uuid;primaryKey"`
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

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
