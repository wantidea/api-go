package blog

type Setting struct {
	ID        int    `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	UpdatedAt int    `json:"updated_at"`
}

func (m *Setting) TableName() string {
	return "blog_setting"
}
