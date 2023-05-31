package user_validation

type Update struct {
	Name string `form:"name" binding:"required,min=3"`
}
