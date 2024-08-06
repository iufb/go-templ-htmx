package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(user User) error
}
type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}
type RegisterUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required,min=6,max=12"`
}
