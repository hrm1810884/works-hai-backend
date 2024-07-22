package impl_repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hrm1810884/works-hai-backend/config"
	"github.com/hrm1810884/works-hai-backend/domain/entity/user"
	impl_repository "github.com/hrm1810884/works-hai-backend/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestImplUserRepository_Integration(t *testing.T) {
	t.Parallel()
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

	userRepo, err := impl_repository.NewImplUserRepository(ctx)
	if err != nil {
		t.Fatalf("failed to get user repo: %v", err)
	}

	// テストデータの作成
	id := uuid.New()
	userId, err := user.NewUserId(id)
	if err != nil {
		t.Fatalf("failed to get user id: %v", err)
	}
	userData := user.NewUser(
		*userId,
		*user.NewPosition(0, 0),
		"hogehoge",
		time.Now(),
		time.Now(),
	)

	// Createのテスト
	err = userRepo.Create(*userData)
	assert.NoError(t, err, "failed to create user")

	// FindByIdのテスト
	foundUser, err := userRepo.FindById(*userData.GetId())
	assert.NoError(t, err, "failed to find user by id")
	assert.Equal(t, userData.GetId().ToId(), foundUser.GetId().ToId(), "user id does not match")
	assert.Equal(t, userData.GetPosition().GetX(), foundUser.GetPosition().GetX(), "position X does not match")
	assert.Equal(t, userData.GetPosition().GetY(), foundUser.GetPosition().GetY(), "position Y does not match")

	foundUser, err = userRepo.FindById(*userData.GetId())
	assert.NoError(t, err, "failed to find user by id after update")
	assert.Equal(t, userData.GetPosition().GetX(), foundUser.GetPosition().GetX(), "updated position X does not match")
	assert.Equal(t, userData.GetPosition().GetY(), foundUser.GetPosition().GetY(), "updated position Y does not match")

	// Deleteのテスト
	// err = userRepo.Delete(*userData.GetId())
	// assert.NoError(t, err, "failed to delete user")

	// foundUser, err = userRepo.FindById(*userData.GetId())
	// assert.Error(t, err, "expected error when finding deleted user")
	// assert.Nil(t, foundUser, "expected no user after deletion")
}
