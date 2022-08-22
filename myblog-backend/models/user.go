package models

type IdUser = uint

type User struct {
	Id       IdUser `gorm:"primarykey"`
	Name     string `gorm:"type:varchar(100);not null"`
	LastName string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null"`
	DBModel
}

type UserCreate struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserResponse struct {
	Id       IdUser `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

func NewUser(data UserCreate) *User {
	return &User{
		Name:     data.Name,
		LastName: data.LastName,
		Password: data.Password,
		Email:    data.Email,
	}
}
