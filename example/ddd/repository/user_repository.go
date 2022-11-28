package repository

import (
	"fmt"
	"github.com/shukubota/go-api-template/example/ddd/domain/entity"
	"github.com/shukubota/go-api-template/example/ddd/domain/repository_interface"
)

type UserRepository struct{}

func NewRegisterUserRepository() repository_interface.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Find(id uint) (*entity.User, error) {
	// data := dbから取得(SELECT id, name FROM users WHERE id = :id;)
	return &entity.User{
		ID: &id,
	}, nil
}

func (r *UserRepository) Save(u *entity.User) error {
	if u.ID != nil {
		// dbにupdate処理 (UPDATE users SET name = u.Name WHERE id = u.ID;)
	} else {
		// dbにinsert処理 (INSERT INTO users(id, name) VALUES(u.ID, u.Name);)
	}
	fmt.Println(u)
	return nil
}
