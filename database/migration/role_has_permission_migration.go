package migration

type RoleHasPermission struct {
	Id int `json:"id" gorm:"primary_key"`

	RoleId int
	Role   Role `json:"role_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	PermissionId int
	Permission   Permission `json:"permission_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt string `json:"created_at" gorm:"type:timestamp(0)"`
	UpdatedAt string `json:"updated_at" gorm:"type:timestamp(0)"`
}
