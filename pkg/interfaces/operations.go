package interfaces

import (
	"gonum.org/v1/gonum/mat"
)
type Geometry interface {
	Dimension() int
	CrossProduct(*mat.VecDense, *mat.VecDense) *mat.VecDense
}
