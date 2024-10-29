package usecase

import (
	"github.com/hrm1810884/works-hai-backend/domain/entity/history"
	"github.com/hrm1810884/works-hai-backend/domain/repository"
)

type RebootHistoryUsecase struct {
	historyRepository repository.HistoryRepository
}

func NewRebootHistoryUsecase(historyRepository repository.HistoryRepository) *RebootHistoryUsecase {
	return &RebootHistoryUsecase{historyRepository}
}

func (u *RebootHistoryUsecase) RebootHistory() (*history.History, error) {
	latestHistory, err := u.historyRepository.FindLatest()
	if err != nil {
		return nil, err
	}
	nextVersion := latestHistory.Version.GetNextVersion()
	newHistory := history.NewHistory(nextVersion.GetVersion())
	err = u.historyRepository.Create(*newHistory)
	if err != nil {
		return nil, err
	}
	return newHistory, nil
}
