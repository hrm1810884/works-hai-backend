package impl_repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/hrm1810884/works-hai-backend/config"
	"github.com/hrm1810884/works-hai-backend/domain/entity/history"
	"github.com/hrm1810884/works-hai-backend/infrastructure/impl/database"
)

type ImplHistoryRepository struct {
	Client *firestore.Client
}

func NewImplHistoryRepository(ctx context.Context) (*ImplHistoryRepository, error) {
	app, err := config.InitializeApp()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	return &ImplHistoryRepository{Client: client}, nil
}

func (hr *ImplHistoryRepository) Create(history history.History) error {
	ctx := context.Background()
	historyData := ConvertHistoryToData(history)

	err := database.CreateHistory(hr.Client, ctx, historyData)
	if err != nil {
		return err
	}

	return nil
}

func (hr *ImplHistoryRepository) FindByVersion(version int) (*history.History, error) {
	ctx := context.Background()

	historyData, err := database.FindByVersion(hr.Client, ctx, version)
	if err != nil {
		return nil, err
	}

	return ConvertDataToHistory(*historyData)
}

func (hr *ImplHistoryRepository) FindLatest() (*history.History, error) {
	ctx := context.Background()

	historyData, err := database.FindLatestVersion(hr.Client, ctx)
	if err != nil {
		return nil, err
	}

	return ConvertDataToHistory(*historyData)
}

func ConvertDataToHistory(data database.HistoryData) (*history.History, error) {
	return &history.History{
		HistoryId: data.HistoryId,
		Version:   *history.NewVersion(data.Version),
	}, nil
}

func ConvertHistoryToData(history history.History) database.HistoryData {
	return database.HistoryData{
		HistoryId: history.HistoryId,
		Version:   history.Version.GetVersion(),
	}
}
