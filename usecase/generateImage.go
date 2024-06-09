package usecase

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/hrm1810884/works-hai-backend/entity"
	"github.com/hrm1810884/works-hai-backend/repository"
	"github.com/hrm1810884/works-hai-backend/usecase/service"
)

type IGenerateImage interface {
	GenerateAIDrawing(ctx context.Context) (string, error)
}

type GenerateImageUsecase struct {
	QuadImages entity.QuadImagesEntity
}

func NewGenerateImageUsecase(ctx context.Context, currentPosition entity.IPosition) IGenerateImage {
	quadImages := entity.NewQuadImages(ctx, currentPosition)
	return &GenerateImageUsecase{
		QuadImages: quadImages,
	}
}

func (u *GenerateImageUsecase) GenerateAIDrawing(ctx context.Context) (string, error) {
	imagePaths := []string{}

	s, err := service.NewGetSignedUrlService(ctx)
	if err != nil {
		return "", err
	}

	for pos, cfg := range u.QuadImages.Config {
		if cfg.IsDrawn {
			signedUrl, err := s.GetSignedUrl(cfg.ResourceName, "GET")
			if err != nil {
				return "", err
			}

			err = repository.DownloadImage(signedUrl, cfg.SavedPath)
			if err != nil {
				return "", err
			}

			imagePaths = append(imagePaths, fmt.Sprintf("%s=%s", pos, cfg.SavedPath))

		}
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
