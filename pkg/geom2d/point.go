package geom2d

import (
	"gonum.org/v1/gonum/mat"
)

type Point2D struct {
	*mat.VecDense
}

func NewPoint(x, y float64) *Point2D {
	point := mat.NewVecDense(1, []float64{x, y})
	return &Point2D{point}
}
