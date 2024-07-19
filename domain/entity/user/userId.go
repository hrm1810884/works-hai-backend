package user

import "fmt"

type UserId struct {
	id string
}

func NewUserId(id string) (*UserId, error) {
	userId := new(UserId)
	if id == "" {
		return nil, fmt.Errorf("NewUserId Error: userId is required")
	}
	userId.id = id
	return userId, nil
}

func (u *UserId) GetDrawingNameFromId() string {
	return u.id + ".png"
}

func (u *UserId) ToId() string {
	return u.id
}
