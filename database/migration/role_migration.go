package migration

type Role struct {
	Id int `json:"id" gorm:"primary_key"`

	RoleHasPermissions []RoleHasPermission

	Name string `json:"name" gorm:"type:varchar;not_null;unique"`

	CreatedAt string `json:"created_at" gorm:"type:timestamp(0)"`
	UpdatedAt string `json:"updated_at" gorm:"type:timestamp(0)"`
}
