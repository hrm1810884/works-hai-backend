package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/hrm1810884/works-hai-backend/domain"
	"google.golang.org/api/iterator"
)

type ImplUserRepository struct {
	Client *firestore.Client
}

type UserData struct {
	UserId    string
	PosX      int
	PosY      int
	Url       string
	IsDrawn   bool
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(client *firestore.Client, ctx context.Context, historyId string, user UserData) error {
	_, err := client.Collection("history").Doc(historyId).Collection("users").Doc(user.UserId).Set(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to add user to Firestore: %w", err)
	}

	return nil
}

func FindById(client *firestore.Client, ctx context.Context, historyId string, userId string) (*UserData, error) {
	query := client.Collection("history").Doc(historyId).Collection("users").
		Where("UserId", "==", userId).
		OrderBy("UpdatedAt", firestore.Desc).
		Limit(1)

	doc, err := query.Documents(ctx).Next()
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	userData := UserData{}
	if err := doc.DataTo(&userData); err != nil {
		return nil, fmt.Errorf("failed to convert Firestore document to UserData: %w", err)
	}

	return &userData, nil
}

func FindByPos(client *firestore.Client, ctx context.Context, historyId string, posX int, posY int) (*UserData, error) {
	query := client.Collection("history").Doc(historyId).Collection("users").
		Where("PosX", "==", posX).
		Where("PosY", "==", posY).
		OrderBy("UpdatedAt", firestore.Desc).
		Limit(1)

	doc, err := query.Documents(ctx).Next()
	if errors.Is(err, iterator.Done) {
		return nil, domain.ErrNoLatestUser
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	userData := UserData{}
	if err := doc.DataTo(&userData); err != nil {
		return nil, fmt.Errorf("failed to convert Firestore document to UserData: %w", err)
	}

	return &userData, nil
}

func FindLatest(client *firestore.Client, ctx context.Context, historyId string) (*UserData, error) {
	query := client.Collection("history").Doc(historyId).Collection("users").
		OrderBy("UpdatedAt", firestore.Desc).
		Limit(1)

	doc, err := query.Documents(ctx).Next()
	if errors.Is(err, iterator.Done) {
		return nil, domain.ErrNoLatestUser
	} else if err != nil {
		return nil, fmt.Errorf("failed to get latest user: %w", err)
	}

	userData := UserData{}
	if err := doc.DataTo(&userData); err != nil {
		return nil, fmt.Errorf("failed to convert Firestore document to UserData: %w", err)
	}

	return &userData, err
}

func Update(client *firestore.Client, ctx context.Context, historyId string, user UserData) error {
	_, err := client.Collection("history").Doc(historyId).Collection("users").Doc(user.UserId).Set(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to update user in Firestore: %w", err)
	}

	return nil
}

func Delete(client *firestore.Client, ctx context.Context, historyId string, userId string) error {
	_, err := client.Collection("history").Doc(historyId).Collection("users").Doc(userId).Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete user from Firestore: %w", err)
	}

	return nil
}
