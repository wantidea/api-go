package user

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	RoleId    int64  `json:"role_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
	State     int8   `json:"state"`
}

func (m *User) TableName() string {
	return "user_info"
}
