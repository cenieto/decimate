package geom2d

import (
	"gonum.org/v1/gonum/mat"
)

// Point2D represents a 2D point using a vector with two components (x, y).
type Point2D struct {
	*mat.VecDense // Embeds VecDense for matrix operations
}

// NewPoint creates a new Point2D instance.
//
// Arguments:
//   x (float64): The x-coordinate of the point.
//   y (float64): The y-coordinate of the point.
//
// Returns:
//   *Point2D: A pointer to the newly created Point2D, which contains a 2D vector.
func NewPoint(x, y float64) *Point2D {
	point := mat.NewVecDense(2, []float64{x, y})
	return &Point2D{point}
}
