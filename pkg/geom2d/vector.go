package geom2d

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// Vector2D is a wrapper around *mat.VecDense that only represents 2D vectors.
// It provides a more semantic representation of a 2D vector in the geometry package, ensuring
// that the vectors are always 2D in nature and allows operations on them.
type Vector2D struct {
	*mat.VecDense
}

// NewVector creates and returns a new 2D vector as a Vector2D.
//
// Parameters:
//   - x (float64): The x-coordinate (horizontal component) of the vector.
//   - y (float64): The y-coordinate (vertical component) of the vector.
//
// Returns:
//
//	*Vector2D: A pointer to a 2D vector represented as a Vector2D.
//
// Example:
//
//	v := geom2d.NewVector(3.0, 4.0)
//	fmt.Println(v)  // Output: [3.0, 4.0]
func NewVector(x, y float64) *Vector2D {
	return &Vector2D{
		VecDense: mat.NewVecDense(2, []float64{x, y}),
	}
}

// NewVectorTwoPoints creates a 2D vector from two points.
//
// The resulting vector is calculated as the difference between the coordinates of the second
// point (p2) and the first point (p1): p2 - p1.
//
// Parameters:
//   - p1 (*Point2D): The starting point of the vector.
//   - p2 (*Point2D): The ending point of the vector.
//
// Returns:
//
//	*Vector2D: A pointer to a 2D vector represented as a Vector2D.
//
// Example:
//
//	p1 := geom2d.NewPoint(0.0, 1.0)
//	p2 := geom2d.NewPoint(1.0, 0.0)
//	v := geom2d.NewVectorTwoPoints(p1, p2)
//	fmt.Println(v)  // Output: [1.0, -1.0]
func NewVectorTwoPoints(p1, p2 *Point2D) *Vector2D {
	coordinates := []float64{
		p2.At(0, 0) - p1.At(0, 0),
		p2.At(1, 0) - p1.At(1, 0),
	}
	vec := mat.NewVecDense(2, coordinates) // Create a vector with the computed coordinates

	return &Vector2D{
		VecDense: vec,
	}
}

// String returns a string representation of the Vector object.
//
// Returns:
//
//	string: A string representation of the Vector object.
func (v Vector2D) String() string {
	return fmt.Sprintf("%v", mat.Formatted(v.VecDense))
}
