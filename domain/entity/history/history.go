package history

import "github.com/hrm1810884/works-hai-backend/domain/entity/user"

type History struct {
	HistoryId int
	Version   int
	Data      user.User
}

func NewHistory(historyId int, version int, data user.User) *History {
	return &History{
		HistoryId: historyId,
		Version:   version,
		Data:      data,
	}
}
