package usecase

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hrm1810884/works-hai-backend/repository"
)

type IGenerateImage interface {
	GenerateAIDrawing() (string, error)
}

func (u *GetSignedUrlsUsecase) GenerateAIDrawing(ctx context.Context) (string, error) {
	imagePaths := []string{}

	signedUrls, err := u.GenerateSignedURLs(ctx, "GET")
	if err != nil {
		return "", err
	}

	if u.isDrawn.Top {
		err := repository.DownloadImage(signedUrls.Top, "./image/Top.png")
		if err != nil {
			return "", fmt.Errorf("top image generation: %w", err)
		}
		imagePaths = append(imagePaths, "./image/Top.png")
	}
	if u.isDrawn.Right {
		err := repository.DownloadImage(signedUrls.Right, "./image/Left.png")
		if err != nil {
			return "", fmt.Errorf("right image generation: %w", err)
		}
		imagePaths = append(imagePaths, "./image/Right.png")
	}
	if u.isDrawn.Bottom {
		log.Printf("%q", signedUrls.Bottom)
		err := repository.DownloadImage(signedUrls.Bottom, "./image/Bottom.png")
		if err != nil {
			return "", fmt.Errorf("bottom image generation: %w", err)
		}
		imagePaths = append(imagePaths, "./image/Bottom.png")
	}
	if u.isDrawn.Left {
		err := repository.DownloadImage(signedUrls.Left, "./image/Left.png")
		if err != nil {
			return "", fmt.Errorf("left image generation: %w", err)
		}
		imagePaths = append(imagePaths, "./image/Left.png")
	}

	err = executePythonScript("./usecase/scripts/process_image.py", imagePaths)
	if err != nil {
		return "", fmt.Errorf("failed to execute Python script: %w", err)
	}

	return "Images processed successfully", nil
}

func executePythonScript(scriptPath string, imagePaths []string) error {
	cmd := exec.Command("python3", append([]string{scriptPath}, imagePaths...)...) //nolint:gosec
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
