package models

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string 
}

func (u User) Exists() bool {
	return u.ID != 0 && u.Username != ""
}