package models

import (
	appService "gRPC/app_services/internal/services"
	"time"
)

type User struct {
	UUID string `gorm:"primaryKey" json:"uuid"`

	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	DOB int64 `json:"dob"`
	Email string `json:"email"`

	// authenticate_services
	Username string `json:"username"`
	Password string `json:"password"`

	Platforms []Platform `json:"platforms" gorm:"many2many:user_platforms"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Map(user *appService.UserRequest) {
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.DOB = user.Dob
	u.Email = user.Email
	u.Username = user.Username
	u.Password = user.Password
}
