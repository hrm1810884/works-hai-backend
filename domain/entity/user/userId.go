package user

import (
	"fmt"

	"github.com/google/uuid"
)

type UserId struct {
	id uuid.UUID
}

func NewUserId(id uuid.UUID) (*UserId, error) {
	userId := new(UserId)
	if id == uuid.Nil {
		return nil, fmt.Errorf("NewUserId Error: userId is required")
	}
	userId.id = id
	return userId, nil
}

func (u *UserId) GetDrawingName() string {
	return u.id.String() + ".png"
}

func (u *UserId) ToId() string {
	return u.id.String()
}
