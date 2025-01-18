package geom3d

import (
	"gonum.org/v1/gonum/mat"
)

// Line3D represents a line in 3D space defined by two points (Point1 and Point2).
type Line3D struct {
	Point1 *Point3D // Starting point of the line
	Point2 *Point3D // Ending point of the line
}

// NewLine creates a new Line3D instance.
//
// Arguments:
//
//	pa (*Point3D): The first point defining the line.
//	pb (*Point3D): The second point defining the line.
//
// Returns:
//
//	*Line3D: A pointer to the newly created Line3D object.
func NewLine(pa, pb *Point3D) *Line3D {
	line := Line3D{Point1: pa, Point2: pb}
	return &line
}

// VectorDirector calculates the direction vector of the line.
//
// The direction vector is computed as the difference between Point2 and Point1.
//
// Returns:
//
//	*Vector3D: A pointer to a Vector3D representing the direction of the line.
func (l Line3D) VectorDirector() *Vector3D {
	var result mat.VecDense
	result.SubVec(l.Point2, l.Point1) // Calculate Point2 - Point1
	vector := NewVector(result.At(0, 0), result.At(1, 0), result.At(2, 0))
	return vector
}
