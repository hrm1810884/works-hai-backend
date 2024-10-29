package repository

import "github.com/hrm1810884/works-hai-backend/domain/entity/history"

type HistoryRepository interface {
	Create(history history.History) error
	FindLatest() (*history.History, error)
	FindByVersion(version int) (*history.History, error)
}
