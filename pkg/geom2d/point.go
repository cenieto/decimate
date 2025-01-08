package geom2d

type Point2D struct {
	X, Y float64
}

func NewPoint(x, y float64) Point2D {
	return Point2D{X: x, Y: y}
}
