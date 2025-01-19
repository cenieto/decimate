package interfaces

import (
	"gonum.org/v1/gonum/mat"
)

type Geometry interface {
	Dimension() int
	CrossProduct(*mat.VecDense, *mat.VecDense) *mat.VecDense
	CrossProductNorm(*mat.VecDense, *mat.VecDense) float64
	DoubleAreaTriangle(*mat.VecDense, *mat.VecDense, *mat.VecDense) float64
	DistancePointLine(*mat.VecDense, *mat.VecDense, *mat.VecDense) float64
	NewPoint(float64, float64) *mat.VecDense
}

type Point interface {
	String() string
	Dimension() int
	At(int, int) float64
}
