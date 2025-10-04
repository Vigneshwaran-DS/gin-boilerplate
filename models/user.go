package models

type User struct {
	BaseModel
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	FullName string `json:"full_name"`
}

func (User) TableName() string {
	return "users"
}
