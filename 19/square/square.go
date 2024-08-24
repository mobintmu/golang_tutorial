package square

import "golang_tutorial/19/point"

type Square struct {
	X [4]point.Point
}

func NewSquare(x1, y1, x2, y2, x3, y3, x4, y4 float32) *Square {
	return &Square{[4]point.Point{*point.NewPoint(x1, y1), *point.NewPoint(x2, y2), *point.NewPoint(x3, y3), *point.NewPoint(x4, y4)}}
}

func (s *Square) Get() [4]point.Point {
	return s.X
}
