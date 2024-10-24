package repository

import (
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
)

type UserRepository interface {
	Create(user user.User) error
	FindById(userId user.UserId) (*user.User, error)
	FindByPos(pos user.Position) (*user.User, error)
	FindLatest() (*user.User, error)
	GetLatestArray() ([]user.User, error)
	Update(user user.User) error
	Delete(userId user.UserId) error
	DeleteImagesExceptCenter() (error)
}
