package usecase

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hrm1810884/works-hai-backend/entity"
	"github.com/hrm1810884/works-hai-backend/usecase/service"
)

type IGetSignedUrls interface {
	GenerateSignedUrls() (string, error)
}

type IsDrawn struct {
	Top    bool `json:"top"`
	Right  bool `json:"right"`
	Bottom bool `json:"bottom"`
	Left   bool `json:"left"`
}

type SignedUrls struct {
	Top    string `json:"top"`
	Right  string `json:"right"`
	Bottom string `json:"bottom"`
	Left   string `json:"left"`
}

type GetSignedUrlsUsecase struct {
	currentPosition entity.IPosition
	isDrawn         IsDrawn
}

func NewGetSignedUrlsUsecase(prevPositon entity.IPosition) (*GetSignedUrlsUsecase, error) {
	currentPosition, err := prevPositon.GetNext()
	if err != nil {
		return nil, err
	}
	return &GetSignedUrlsUsecase{
		currentPosition: currentPosition,
		isDrawn: IsDrawn{
			Top:    false,
			Right:  false,
			Bottom: true,
			Left:   true,
		},
	}, nil
}

func (u *GetSignedUrlsUsecase) GenerateSignedURLs(ctx context.Context, method string) (*SignedUrls, error) {
	getSignedUrlService, err := service.NewGetSignedUrlService(ctx)
	if err != nil {
		return nil, err
	}

	isDrawnValue := reflect.ValueOf(u.isDrawn)
	isDrawnType := isDrawnValue.Type()

	signedUrls := &SignedUrls{}

	for i := 0; i < isDrawnValue.NumField(); i++ {
		field := isDrawnType.Field(i)
		flag := isDrawnValue.Field(i).Bool()
		fileName := ""

		switch field.Name {
		case "Top":
			fileName = fmt.Sprintf("%v_%v.png", u.currentPosition.GetX(), u.currentPosition.GetY()+1)
		case "Right":
			fileName = fmt.Sprintf("%v_%v.png", u.currentPosition.GetX()+1, u.currentPosition.GetY())
		case "Bottom":
			fileName = fmt.Sprintf("%v_%v.png", u.currentPosition.GetX(), u.currentPosition.GetY()-1)
		case "Left":
			fileName = fmt.Sprintf("%v_%v.png", u.currentPosition.GetX()-1, u.currentPosition.GetY())
		}

		if flag {
			url, err := getSignedUrlService.GetSignedUrl(fileName, method)
			if err != nil {
				return nil, fmt.Errorf("failed to generate signed URL for %s: %w", fileName, err)
			}

			switch field.Name {
			case "Top":
				signedUrls.Top = url
			case "Right":
				signedUrls.Right = url
			case "Bottom":
				signedUrls.Bottom = url
			case "Left":
				signedUrls.Left = url
			}
		}
	}

	return signedUrls, nil
}
