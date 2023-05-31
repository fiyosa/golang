package auth_validation

type Login struct {
	Username string `form:"username" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}
