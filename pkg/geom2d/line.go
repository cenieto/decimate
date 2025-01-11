package geom2d

import (
	"gonum.org/v1/gonum/mat"
)

type Line2D struct {
	PointA *Point2D
	PointB *Point2D
}

func NewLine(pa, pb *Point2D) *Line2D {
	line := Line2D{PointA: pa, PointB: pb}
	return &line
}

func (l Line2D) VectorDirector() *Vector2D {
	var result mat.VecDense
	result.SubVec(l.PointB, l.PointA)
	vector := NewVector(result.At(0, 0), result.At(1, 0))
	return vector
}
