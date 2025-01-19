package geom3d

import (
	"fmt"
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
	return &Vector3D{
		VecDense: mat.NewVecDense(3, []float64{x, y, z}),
	}
}

// NewVectorTwoPoints creates a 3D vector from two points.
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
//	*Vector3D: A pointer to a 3D vector represented as a Vector3D.
//
// Example:
//
//	p1 := geom3d.NewPoint(0.0, 1.0)
//	p2 := geom3d.NewPoint(1.0, 0.0)
//	v := geom3d.NewVectorTwoPoints(p1, p2)
//	fmt.Println(v)  // Output: [1.0, -1.0]
func NewVectorTwoPoints(p1, p2 *Point3D) *Vector3D {
	coordinates := []float64{
		p2.At(0, 0) - p1.At(0, 0),
		p2.At(1, 0) - p1.At(1, 0),
		p2.At(2, 0) - p1.At(2, 0),
	}
	return &Vector3D{
		VecDense: mat.NewVecDense(3, coordinates),
	}
}

// String returns a string representation of the Vector object.
//
// Returns:
//
//	string: A string representation of the Vector object.
func (v Vector3D) String() string {
	return fmt.Sprintf("%v", mat.Formatted(v.VecDense))
}

// Length calculates the length of the vector.
//
// The length is computed as the norm of the vector.
//
// Returns:
//
//	float64: The length of the vector.
func (v Vector3D) Length() float64 {
	return v.Norm(2)
}
