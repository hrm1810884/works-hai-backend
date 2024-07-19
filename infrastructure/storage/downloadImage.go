package storage

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func validateURL(u string) (*url.URL, error) {
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, err
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return nil, fmt.Errorf("invalid URL scheme")
	}

	return parsedURL, nil
}

func DownloadImage(url, filePath string) error {
	validUrl, err := validateURL(url)
	if err != nil {
		// Handle the error appropriately
		return fmt.Errorf("invalid URL: %w", err)
	}

	resp, err := http.Get(validUrl.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Print("%w", url)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image: received non-200 response code %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0600)
}
