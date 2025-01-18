package geom3d

import (
	"gonum.org/v1/gonum/mat"
)

// Point3D represents a 3D point using a vector with two components (x, y).
type Point3D struct {
	*mat.VecDense // Embeds VecDense for matrix operations
}

// NewPoint creates a new Point3D instance.
//
// Arguments:
//
//	x (float64): The x-coordinate of the point.
//	y (float64): The y-coordinate of the point.
//
// Returns:
//
//	*Point3D: A pointer to the newly created Point3D, which contains a 3D vector.
func NewPoint(x, y, z float64) *Point3D {
	point := mat.NewVecDense(3, []float64{x, y, z})
	return &Point3D{point}
}
