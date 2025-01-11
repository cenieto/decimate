package geom3d

import (
	"gonum.org/v1/gonum/mat"
)

// Vector3D is a wrapper around *mat.VecDense that only represents 3D vectors.
// It provides a more semantic representation of a 3D vector in the geometry package, ensuring
// that the vectors are always 3D in nature and allows operations on them.
type Vector3D struct {
	*mat.VecDense
}

// NewVector creates and returns a new 3D vector as a Vector3D.
// The function takes two floating-point values that represent the x and y components of the vector.
// The returned Vector3D is a wrapper around the gonum *mat.VecDense type, initialized with the
// provided x and y values.
//
// Parameters:
//   - x: The x-coordinate (horizontal component) of the vector.
//   - y: The y-coordinate (vertical component) of the vector.
//
// Returns:
//   - Vector3D: A 3D vector represented as a Vector3D, which is a wrapper around *mat.VecDense.
//
// Example usage:
//
//	v := geom3D.NewVector(3.0, 4.0)
//	fmt.Println(v)  // Output: [3.0, 4.0]
func NewVector(x, y, z float64) *Vector3D {
	// Create a 3D vector using gonum's NewVecDense
	vec := mat.NewVecDense(3, []float64{x, y, z})
	return &Vector3D{vec}
}
