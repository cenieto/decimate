package geom2d

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// Point2D represents a 2D point using a vector with two components (x, y).
type Point2D struct {
	*mat.VecDense
}

// NewPoint creates a new Point2D instance.
//
// Arguments:
//
//	x (float64): The x-coordinate of the point.
//	y (float64): The y-coordinate of the point.
//
// Returns:
//
//	*Point2D: A pointer to the newly created Point2D, which contains a 2D vector.
func NewPoint(x, y float64) *Point2D {
	return &Point2D{
		VecDense: mat.NewVecDense(2, []float64{x, y}),
	}
}

// String returns a string representation of the Point object.
//
// Returns:
//
//	string: A string representation of the Point2D object.
func (p Point2D) String() string {
	return fmt.Sprintf("Point{%v}", mat.Formatted(p.VecDense))
}
