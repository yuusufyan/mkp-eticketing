package models

func All() []interface{} {
	return []interface{}{
		&Role{},
		&User{},
		&Permission{},
		&UsersRoles{},
		&RolesPermissions{},
		&Menu{},
		&PermissionMenu{},
	}
}
