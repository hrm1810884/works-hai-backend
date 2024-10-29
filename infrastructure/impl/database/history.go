package database

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
)

type ImplHistoryRepository struct {
	Client *firestore.Client
}

type HistoryData struct {
	HistoryId string
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateHistory(client *firestore.Client, ctx context.Context, history HistoryData) error {
	_, err := client.Collection("history").Doc(history.HistoryId).Set(ctx, history)
	if err != nil {
		return fmt.Errorf("failed to add history to Firestore: %w", err)
	}

	return nil
}

func FindByVersion(client *firestore.Client, ctx context.Context, version int) (*HistoryData, error) {
	query := client.Collection("history").
		Where("Version", "==", version).
		Limit(1)

	doc, err := query.Documents(ctx).Next()
	if err != nil {
		return nil, fmt.Errorf("failed to get history: %w", err)
	}

	historyData := HistoryData{}
	if err := doc.DataTo(&historyData); err != nil {
		return nil, fmt.Errorf("failed to convert Firestore document to HistoryData: %w", err)
	}

	return &historyData, nil
}

func FindLatestVersion(client *firestore.Client, ctx context.Context) (*HistoryData, error) {
	query := client.Collection("history").
		OrderBy("Version", firestore.Desc).
		Limit(1)

	doc, err := query.Documents(ctx).Next()
	if err != nil {
		return nil, fmt.Errorf("failed to get history: %w", err)
	}

	historyData := HistoryData{}
	if err := doc.DataTo(&historyData); err != nil {
		return nil, fmt.Errorf("failed to convert Firestore document to HistoryData: %w", err)
	}

	return &historyData, nil
}
