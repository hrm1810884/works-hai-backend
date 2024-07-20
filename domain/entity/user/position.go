package user

import (
	"fmt"
	"math"
)

type Position struct {
	X int
	Y int
}

func NewPosition(x int, y int) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func (p *Position) GetNext() (*Position, error) {
	loopNum := p.GetLoopNum()

	switch {
	case isOnBottomSide(p.X, p.Y, loopNum):
		return &Position{
			X: p.X + 1,
			Y: p.Y,
		}, nil
	case isOnRightSide(p.X, p.Y, loopNum):
		return &Position{
			X: p.X,
			Y: p.Y + 1,
		}, nil
	case isOnTopSide(p.X, p.Y, loopNum):
		return &Position{
			X: p.X - 1,
			Y: p.Y,
		}, nil
	case isOnLeftSide(p.X, p.Y, loopNum):
		return &Position{
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

// nolint:unparam
func isOnLeftSide(x int, y int, loopNum int) bool {
	// y is not used in this function, but we want to keep the signature consistent
	return x == -1*loopNum
}
func isOnBottomSide(x int, y int, loopNum int) bool {
	return -1*loopNum < x && x < loopNum && y == -1*loopNum
}

func (p *Position) GetX() int {
	return p.X
}

func (p *Position) GetY() int {
	return p.Y
}

func (p *Position) GetLoopNum() int {
	return max(abs(p.X), abs(p.Y))
}

func (p *Position) GetTop() *Position {
	loopNum := p.GetLoopNum()
	x := p.GetX()
	y := p.GetY()
	isDrawn := -1*loopNum <= x && x < loopNum && y == -1*loopNum || x == -1*loopNum && -1*loopNum <= y && y < loopNum
	if isDrawn {
		return NewPosition(x, y)
	} else {
		return nil
	}
}

func (p *Position) GetLeft() *Position {
	loopNum := p.GetLoopNum()
	x := p.GetX()
	y := p.GetY()
	isDrawn := -1*(loopNum-1) < x && x <= loopNum && y == -1*loopNum || x == loopNum && -1*loopNum <= y && y <= loopNum-1
	if isDrawn {
		return NewPosition(x, y)
	} else {
		return nil
	}
}

func (p *Position) GetBottom() *Position {
	loopNum := p.GetLoopNum()
	x := p.GetX()
	y := p.GetY()
	isDrawn := -1*loopNum < x && x <= loopNum && y == loopNum || x == loopNum && -1*loopNum < y && y <= loopNum
	if isDrawn {
		return NewPosition(x, y)
	} else {
		return nil
	}
}

func (p *Position) GetRight() *Position {
	loopNum := p.GetLoopNum()
	x := p.GetX()
	y := p.GetY()
	isDrawn := -1*loopNum <= x && x < loopNum && y == loopNum || x == -1*loopNum && -1*loopNum <= y && y <= loopNum
	if isDrawn {
		return NewPosition(x, y)
	} else {
		return nil
	}
}
