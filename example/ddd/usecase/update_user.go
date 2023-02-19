package usecase

import (
	"github.com/shukubota/go-api-template/example/ddd/domain/repository_interface"
)

type UpdateUserUseCase struct {
	r repository_interface.UserRepository
}

func NewUpdateUserUseCase(r repository_interface.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		r: r,
	}
}

func (uc *RegisterUserUseCase) Update(id uint, n string) error {
	u, err := uc.r.Find(id)
	if err != nil {
		return err
	}

	err = u.ChangeName(n)
	if err != nil {
		return err
	}

	err = uc.r.Save(u)
	if err != nil {
		return err
	}
	return nil
}
