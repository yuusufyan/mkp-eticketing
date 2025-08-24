package dto

import (
	"auth-rbac/src/utils/enums"

	"github.com/google/uuid"
)

type CreateMenu struct {
	ID       int                `json:"id"`
	Name     string             `json:"name"`
	Status   enums.StatusActive `json:"status"`
	ParentID uuid.UUID          `json:"parent_id"`
}

func NewCreateMenu() CreateMenu {
	return CreateMenu{
		Status: enums.StatusActive_Active, // default value
	}
}

func NewCreateParent() CreateMenu {
	return CreateMenu{
		ParentID: uuid.Nil, // default value
	}
}
