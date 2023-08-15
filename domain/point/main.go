package point

type Point struct {
	x int
	y int
}

var _ PointService = (*Point)(nil)

type PointService interface {
	X() int
	Y() int
	Add(p1, p2 Point) Point
	Sub(p1, p2 Point) Point
	Mul(p1, p2 Point) Point
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) Y() int {
	return p.y
}

func (p *Point) Add(p1, p2 Point) Point {
	return Point{p1.x + p2.x, p1.y + p2.y}
}

func (p *Point) Sub(p1, p2 Point) Point {
	return Point{p1.x - p2.x, p1.y - p2.y}
}

func (p *Point) Mul(p1, p2 Point) Point {
	return Point{p1.x * p2.x, p1.y * p2.y}
}
