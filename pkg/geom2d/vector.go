package geom2d

import (
	"gonum.org/v1/gonum/mat"
)

// Vector2D is a wrapper around *mat.VecDense that only represents 2D vectors.
// It provides a more semantic representation of a 2D vector in the geometry package, ensuring
// that the vectors are always 2D in nature and allows operations on them.
type Vector2D struct {
	*mat.VecDense
}

// NewVector creates and returns a new 2D vector as a Vector2D.
// The function takes two floating-point values that represent the x and y components of the vector.
// The returned Vector2D is a wrapper around the gonum *mat.VecDense type, initialized with the
// provided x and y values.
//
// Parameters:
//   - x: The x-coordinate (horizontal component) of the vector.
//   - y: The y-coordinate (vertical component) of the vector.
//
// Returns:
//   - Vector2D: A 2D vector represented as a Vector2D, which is a wrapper around *mat.VecDense.
//
// Example usage:
//
//	v := geom2d.NewVector(3.0, 4.0)
//	fmt.Println(v)  // Output: [3.0, 4.0]
func NewVector(x, y float64) *Vector2D {
	// Create a 2D vector using gonum's NewVecDense
	vec := mat.NewVecDense(2, []float64{x, y})
	return &Vector2D{vec}
}
