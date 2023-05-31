package migration

type User struct {
	Id int `json:"id" gorm:"primary_key"`

	Oauths       []Oauth
	UserHasRoles []UserHasRole

	Email    string `json:"email" gorm:"type:varchar;not_null;unique"`
	Password string `json:"-" gorm:"type:varchar"`
	Name     string `json:"name" gorm:"type:varchar"`

	CreatedAt string `json:"created_at" gorm:"type:timestamp(0)"`
	UpdatedAt string `json:"updated_at" gorm:"type:timestamp(0)"`
}
