package entity

import "github.com/hrm1810884/works-hai-backend/domain/entity/user"

type Drawing struct {
	FileName  string
	UrlToGet  string
	UrlToPost string
}

func NewDrawing(userId *user.UserId) *Drawing {
	return &Drawing{
		FileName:  userId.GetDrawingNameFromId(),
		UrlToGet:  "",
		UrlToPost: "",
	}
}
