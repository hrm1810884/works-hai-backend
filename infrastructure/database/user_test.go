package database_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/hrm1810884/works-hai-backend/config"
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
	"github.com/hrm1810884/works-hai-backend/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestImplUserRepository_Integration(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	t.Cleanup(func() {
		err := os.Chdir(cwd)
		if err != nil {
			t.Fatalf("failed to get current working directory: %v", err)
		}
	})
	err = os.Chdir("../..")
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	ctx := context.Background()

	// Firestoreクライアントの初期化
	app, err := config.InitializeApp()
	if err != nil {
		t.Fatalf("failed to initialize Firebase app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		t.Fatalf("failed to initialize Firestore client: %v", err)
	}
	defer client.Close()

	userRepo := &database.ImplUserRepository{Client: client}

	// テストデータの作成
	userId, err := user.NewUserId("testUser")
	if err != nil {
		t.Fatalf("failed to get user id: %v", err)
	}
	userData := user.NewUser(
		*userId,
		*user.NewPosition(1, 2),
		time.Now(),
		time.Now(),
	)

	// Createのテスト
	err = userRepo.Create(*userData)
	assert.NoError(t, err, "failed to create user")

	// FindByIdのテスト
	foundUser, err := userRepo.FindById(userData.GetId())
	assert.NoError(t, err, "failed to find user by id")
	assert.Equal(t, userData.GetId().ToId(), foundUser.GetId().ToId(), "user id does not match")
	assert.Equal(t, userData.GetPosition().GetX(), foundUser.GetPosition().GetX(), "position X does not match")
	assert.Equal(t, userData.GetPosition().GetY(), foundUser.GetPosition().GetY(), "position Y does not match")

	foundUser, err = userRepo.FindById(userData.GetId())
	assert.NoError(t, err, "failed to find user by id after update")
	assert.Equal(t, userData.GetPosition().GetX(), foundUser.GetPosition().GetX(), "updated position X does not match")
	assert.Equal(t, userData.GetPosition().GetY(), foundUser.GetPosition().GetY(), "updated position Y does not match")

	// Deleteのテスト
	err = userRepo.Delete(userData.GetId())
	assert.NoError(t, err, "failed to delete user")

	foundUser, err = userRepo.FindById(userData.GetId())
	assert.Error(t, err, "expected error when finding deleted user")
	assert.Nil(t, foundUser, "expected no user after deletion")
}
