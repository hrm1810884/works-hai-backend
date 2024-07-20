package database

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/hrm1810884/works-hai-backend/config"
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
)

type ImplUserRepository struct {
	Client *firestore.Client
}

func NewImplUserRepository(ctx context.Context) (*ImplUserRepository, error) {
	app, err := config.InitializeApp()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	return &ImplUserRepository{Client: client}, nil
}

type UserData struct {
	UserId    string
	PosX      int
	PosY      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ConvertDataToUser(data UserData) (*user.User, error) {
	id, err := uuid.Parse(data.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to convert id to uuid: %w", err)
	}

	userId, err := user.NewUserId(id)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}

	position := user.NewPosition(data.PosX, data.PosY)
	user := user.NewUser(*userId, *position, data.CreatedAt, data.UpdatedAt)
	return user, nil
}

func ConvertUserToData(user user.User) *UserData {
	now := time.Now()
	return &UserData{
		UserId:    user.GetId().ToId(),
		PosX:      user.GetPosition().GetX(),
		PosY:      user.GetPosition().GetY(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: now,
	}
}

func (ur *ImplUserRepository) Create(user user.User) error {

	userData := ConvertUserToData(user)

	ctx := context.Background()
	_, err := ur.Client.Collection("users").Doc(user.GetId().ToId()).Set(ctx, userData)
	if err != nil {
		return fmt.Errorf("failed to add user to Firestore: %w", err)
	}

	return nil
}

func (ur *ImplUserRepository) FindById(userId *user.UserId) (*user.User, error) {
	ctx := context.Background()
	query := ur.Client.Collection("users").
		Where("UserId", "==", userId.ToId()).
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

	return ConvertDataToUser(userData)
}

func (ur *ImplUserRepository) FindByPos(pos user.Position) (*user.User, error) {
	ctx := context.Background()
	query := ur.Client.Collection("users").
		Where("PosX", "==", pos.GetX()).
		Where("PosY", "==", pos.GetY()).
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

	return ConvertDataToUser(userData)
}

func (ur *ImplUserRepository) Update(user *user.User) error {
	userData := ConvertUserToData(*user)
	ctx := context.Background()

	_, err := ur.Client.Collection("users").Doc(user.GetId().ToId()).Set(ctx, userData)
	if err != nil {
		return fmt.Errorf("failed to update user in Firestore: %w", err)
	}

	return nil
}

func (ur *ImplUserRepository) Delete(userId *user.UserId) error {
	ctx := context.Background()
	_, err := ur.Client.Collection("users").Doc(userId.ToId()).Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete user from Firestore: %w", err)
	}

	return nil
}
