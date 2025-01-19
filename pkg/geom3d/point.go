package geom3d

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// Point3D represents a 3D point using a vector with three components (x, y, z).
type Point3D struct {
	*mat.VecDense
}

// NewPoint creates a new Point3D instance.
//
// Arguments:
//
//	x (float64): The x-coordinate of the point.
//	y (float64): The y-coordinate of the point.
//	z (float64): The z-coordinate of the point.
//
// Returns:
//
//	*Point3D: A pointer to the newly created Point3D, which contains a 3D vector.
func NewPoint(x, y, z float64) *Point3D {
	return &Point3D{
		VecDense: mat.NewVecDense(3, []float64{x, y, z}),
	}
}

// String returns a string representation of the Point object.
//
// Returns:
//
//	string: A string representation of the Point2D object.
func (p Point3D) String() string {
	return fmt.Sprintf("Point3D{%v}", mat.Formatted(p.VecDense))
}
