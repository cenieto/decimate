package geom2d

import (
	"gonum.org/v1/gonum/mat"
)

// Line2D represents a line in 2D space defined by two points (Point1 and Point2).
type Line2D struct {
	Point1 *Point2D // Starting point of the line
	Point2 *Point2D // Ending point of the line
}

// NewLine creates a new Line2D instance.
//
// Arguments:
//
//	pa (*Point2D): The first point defining the line.
//	pb (*Point2D): The second point defining the line.
//
// Returns:
//
//	*Line2D: A pointer to the newly created Line2D object.
func NewLine(pa, pb *Point2D) *Line2D {
	line := Line2D{Point1: pa, Point2: pb}
	return &line
}

// VectorDirector calculates the direction vector of the line.
//
// The direction vector is computed as the difference between Point2 and Point1.
//
// Returns:
//
//	*Vector2D: A pointer to a Vector2D representing the direction of the line.
func (l Line2D) VectorDirector() *Vector2D {
	var result mat.VecDense
	result.SubVec(l.Point2, l.Point1) // Calculate Point2 - Point1
	vector := NewVector(result.At(0, 0), result.At(1, 0))
	return vector
}
