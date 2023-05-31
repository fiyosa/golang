package role_validation

type Update struct {
	Name       string   `form:"role" binding:"required,min=3"`
	Permission []string `form:"permission" binding:"required"`
}
