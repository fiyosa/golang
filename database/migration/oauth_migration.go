package migration

type Oauth struct {
	Id int `json:"id" gorm:"primary_key"`

	UserId int
	User   User `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Token   string `json:"token" gorm:"type:varchar"`
	Revoked bool   `json:"revoked" gorm:"type:boolean"`

	CreatedAt string `json:"created_at" gorm:"type:timestamp(0)"`
	UpdatedAt string `json:"updated_at" gorm:"type:timestamp(0)"`
}
