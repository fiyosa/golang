package role_validation

type Store struct {
	Name       string   `form:"role" binding:"required,min=3"`
	Permission []string `form:"permission" binding:"required"`
}
