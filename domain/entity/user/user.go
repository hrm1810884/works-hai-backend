package user

import "time"

type User struct {
	userId    UserId
	position  Position
	url       string
	isDrawn   bool
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(userId UserId, position Position, url string, isDrawn bool, createdAt time.Time, updatedAt time.Time) *User {
	return &User{
		userId:    userId,
		position:  position,
		url:       url,
		isDrawn:   isDrawn,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) GetId() *UserId {
	return &u.userId
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) GetUrl() string {
	return u.url
}

// GetUpdatedAt returns the updatedAt time
func (u *User) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) GetPosition() *Position {
	return &u.position
}

func (u *User) IsDrawn() bool {
	return u.isDrawn
}
