package repository

import (
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
)

type UserRepository interface {
	Save(user *user.User) error
	FindById(userId user.UserId) (*user.User, error)
	FindByPos(pos user.Position) (*user.User, error)
	Update(user *user.User) error
	Delete(userId user.UserId) error
}
