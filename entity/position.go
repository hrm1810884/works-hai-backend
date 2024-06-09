package entity

import (
	"fmt"
	"math"
)

type IPosition interface {
	GetNext() (*PositionEntity, error)
	GetX() int
	GetY() int
	GetLoopNum() int
}

type PositionEntity struct {
	X int
	Y int
}

func NewPositionEntity(x int, y int) IPosition {
	return &PositionEntity{
		X: x,
		Y: y,
	}
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func (p *PositionEntity) GetNext() (*PositionEntity, error) {
	loopNum := p.GetLoopNum()

	switch {
	case isOnBottomSide(p.X, p.Y, loopNum):
		return &PositionEntity{
			X: p.X + 1,
			Y: p.Y,
		}, nil
	case isOnRightSide(p.X, p.Y, loopNum):
		return &PositionEntity{
			X: p.X,
			Y: p.Y + 1,
		}, nil
	case isOnTopSide(p.X, p.Y, loopNum):
		return &PositionEntity{
			X: p.X - 1,
			Y: p.Y,
		}, nil
	case isOnLeftSide(p.X, p.Y, loopNum):
		return &PositionEntity{
			X: p.X,
			Y: p.Y - 1,
		}, nil
	default:
		return nil, fmt.Errorf("invalid position error")
	}
}

func isOnRightSide(x int, y int, loopNum int) bool {
	return x == loopNum && y < loopNum
}

func isOnTopSide(x int, y int, loopNum int) bool {
	return -1*loopNum < x && y == loopNum
}
func isOnLeftSide(x int, y int, loopNum int) bool { //nolint:unparam
	// y is not used in this function, but we want to keep the signature consistent
	return x == -1*loopNum
}
func isOnBottomSide(x int, y int, loopNum int) bool {
	return -1*loopNum < x && x < loopNum && y == -1*loopNum
}

func (p *PositionEntity) GetX() int {
	return p.X
}

func (p *PositionEntity) GetY() int {
	return p.Y
}

func (p *PositionEntity) GetLoopNum() int {
	return max(abs(p.X), abs(p.Y))
}
