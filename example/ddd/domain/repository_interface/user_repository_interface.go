package repository_interface

import "github.com/shukubota/go-api-template/example/ddd/domain/entity"

type UserRepository interface {
	Save(*entity.User) error
	Find(id uint) (*entity.User, error)
}
