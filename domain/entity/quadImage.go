package entity

import (
	"context"
	"fmt"
)

type IGenerateImage interface {
	GenerateAIDrawing() (string, error)
}

type QuadImagesEntity struct {
	Config map[string]struct {
		IsDrawn      bool
		ResourceName string
		SavedPath    string
		Desc         string
	}
}

func NewQuadImages(ctx context.Context, currentPosition IPosition) QuadImagesEntity {
	topFileName := fmt.Sprintf("%v_%v.png", currentPosition.GetX(), currentPosition.GetY()+1)
	rightFileName := fmt.Sprintf("%v_%v.png", currentPosition.GetX()+1, currentPosition.GetY())
	bottomFileName := fmt.Sprintf("%v_%v.png", currentPosition.GetX(), currentPosition.GetY()-1)
	leftFileName := fmt.Sprintf("%v_%v.png", currentPosition.GetX()-1, currentPosition.GetY())

	topImagePath := "./image/Top.png"
	rightImagePath := "./image/Right.png"
	bottomImagePath := "./image/Bottom.png"
	leftImagePath := "./image/Left.png"

	imageConfig := map[string]struct {
		IsDrawn      bool
		ResourceName string
		SavedPath    string
		Desc         string
	}{
		"Top": {
			IsDrawn:      isTopDrawn(currentPosition),
			ResourceName: topFileName,
			SavedPath:    topImagePath,
			Desc:         "top",
		},
		"Right": {
			IsDrawn:      isRightDrawn(currentPosition),
			ResourceName: rightFileName,
			SavedPath:    rightImagePath,
			Desc:         "right",
		},
		"Bottom": {
			IsDrawn:      isBottomDrawn(currentPosition),
			ResourceName: bottomFileName,
			SavedPath:    bottomImagePath,
			Desc:         "bottom",
		},
		"Left": {
			IsDrawn:      isLeftDrawn(currentPosition),
			ResourceName: leftFileName,
			SavedPath:    leftImagePath,
			Desc:         "left",
		},
	}

	return QuadImagesEntity{
		Config: imageConfig,
	}
}

func isTopDrawn(currentPosition IPosition) bool {
	loopNum := currentPosition.GetLoopNum()
	x := currentPosition.GetX()
	y := currentPosition.GetY()
	return -1*loopNum <= x && x < loopNum && y == -1*loopNum || x == -1*loopNum && -1*loopNum <= y && y < loopNum
}

func isLeftDrawn(currentPosition IPosition) bool {
	loopNum := currentPosition.GetLoopNum()
	x := currentPosition.GetX()
	y := currentPosition.GetY()
	return -1*(loopNum-1) < x && x <= loopNum && y == -1*loopNum || x == loopNum && -1*loopNum <= y && y <= loopNum-1
}

func isBottomDrawn(currentPosition IPosition) bool {
	loopNum := currentPosition.GetLoopNum()
	x := currentPosition.GetX()
	y := currentPosition.GetY()
	return -1*loopNum < x && x <= loopNum && y == loopNum || x == loopNum && -1*loopNum < y && y <= loopNum
}
func isRightDrawn(currentPosition IPosition) bool {
	loopNum := currentPosition.GetLoopNum()
	x := currentPosition.GetX()
	y := currentPosition.GetY()
	return -1*loopNum <= x && x < loopNum && y == loopNum || x == -1*loopNum && -1*loopNum <= y && y <= loopNum
}
