package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Replace with your credentials
var clientID = ""
var clientSecret = ""

var config = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	Endpoint:     google.Endpoint,
	RedirectURL:  "http://localhost:8080/callback",
	Scopes:       []string{drive.DriveMetadataReadonlyScope},
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/callback", handleCallback)
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	code := r.URL.Query().Get("code")

	token, err := config.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := config.Client(ctx, token)
	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		http.Error(w, "Unable to retrieve Drive client: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fileList, err := srv.Files.List().PageSize(10).Fields("files(id, name)").Do()
	if err != nil {
		http.Error(w, "Unable to retrieve files: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for _, file := range fileList.Files {
		fmt.Fprintf(w, "Found file: %s (%s)\n", file.Name, file.Id)
	}
}
