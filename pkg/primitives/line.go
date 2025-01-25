package primitives

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// Line represents a line in 2D space defined by two points (Point1 and Point2).
type Line struct {
	Point1 *Point // Starting point of the line
	Point2 *Point // Ending point of the line
}

// NewLine creates a new Line instance.
//
// Arguments:
//
//	pa (*Point): The first point defining the line.
//	pb (*Point): The second point defining the line.
//
// Returns:
//
//	*Line: A pointer to the newly created Line object.
func NewLine(pa, pb *Point) *Line {
	line := Line{Point1: pa, Point2: pb}
	return &line
}

// VectorDirector calculates the direction vector of the line.
//
// The direction vector is computed as the difference between Point2 and Point1.
//
// Returns:
//
//	*Vector2D: A pointer to a Vector2D representing the direction of the line.
func (l Line) VectorDirector() *Vector {
	var result mat.VecDense
	result.SubVec(l.Point2, l.Point1) // Calculate Point2 - Point1
	vector := NewVector([]float64{
		result.At(0, 0), result.At(1, 0),
	})
	return vector
}

// Length calculates the length of the line.
//
// The length is computed as the norm of the direction vector.
//
// Returns:
//
//	float64: The length of the line.
func (l Line) Length() float64 {
	return l.VectorDirector().Length()
}

// String returns a string representation of the Line object.
//
// Returns:
//
//	string: A string representation of the Line object.
func (l Line) String() string {
	return fmt.Sprintf("Line{%v, %v}", l.Point1.String(), l.Point2.String())
}

// Dimension returns the dimension of the line.
//
// Returns:
//
//	int: The dimension of the line.
func (v Line) Dimension() int {
	return 2
}
