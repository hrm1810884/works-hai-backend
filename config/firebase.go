package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func initializeApp() *firebase.App {
	config := &firebase.Config{
		StorageBucket: "storage-exp.appspot.com",
	}
	opt := option.WithCredentialsFile("storage-exp-key.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}
	return app
}
