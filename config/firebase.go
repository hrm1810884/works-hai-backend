package config

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitializeApp() (*firebase.App, error) {
	config := &firebase.Config{
		StorageBucket: "storage-exp.appspot.com",
	}
	opt := option.WithCredentialsFile("config/private/storage-exp-key.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}
