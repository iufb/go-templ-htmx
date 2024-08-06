package user

import (
	"fmt"
	"log"

	"github.com/iufb/go-templ-htmx/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	u := new(types.User)
	if err := s.db.Where("email = ?", email).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User not found")
		}
		return nil, fmt.Errorf("Db error")
	}
	log.Println(u)
	return u, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(user types.User) error {
	res := s.db.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
