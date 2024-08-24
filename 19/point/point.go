package point

type Point struct {
	X float32
	Y float32
}

func NewPoint(x, y float32) *Point {
	return &Point{x, y}
}

func (p *Point) Get() (float32, float32) {
	return p.X, p.Y
}
