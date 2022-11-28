package use_case

import (
	"github.com/shukubota/go-api-template/example/ddd/domain/entity"
	"github.com/shukubota/go-api-template/example/ddd/domain/repository_interface"
)

type RegisterUserUseCase struct {
	r repository_interface.UserRepository
}

func NewRegisterUserUseCase(r repository_interface.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		r: r,
	}
}

func (uc *RegisterUserUseCase) Register(n string) error {
	u := entity.NewUser(nil, n)
	err := uc.r.Save(u)
	if err != nil {
		return err
	}
	return nil
}
