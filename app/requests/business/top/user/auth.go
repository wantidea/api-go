package user

type Auth struct {
	Name     string `form:"name" json:"name" fieldName:"用户名" binding:"required,alphanum,min=3,max=30"`
	Password string `form:"password" fieldName:"密码" json:"password" binding:"required,min=8,max=100"`
}
