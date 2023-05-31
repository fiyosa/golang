package migration

func Models() []interface{} {
	return []interface{}{
		&User{},
		&Oauth{},
		&Role{},
		&Permission{},
		&UserHasRole{},
		&RoleHasPermission{},
	}
}
