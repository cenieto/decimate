package geom2d

type Vector2D struct {
	X, Y float64
}

func NewVector(x, y float64) Vector2D {
	return Vector2D{
		X: x,
		Y: y,
	}
}

func (v Vector2D) Dimension() int {
	return 2
}

func (v Vector2D) Components() []float64 {
	return []float64{v.X, v.Y}
}
