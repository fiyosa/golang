package migration

type UserHasRole struct {
	Id int `json:"id" gorm:"primary_key"`

	UserId int
	User   User `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	RoleId int
	Role   Role `json:"role_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt string `json:"created_at" gorm:"type:timestamp(0)"`
	UpdatedAt string `json:"updated_at" gorm:"type:timestamp(0)"`
}
