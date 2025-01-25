package primitives

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// Vector represents a vector, which is a mathematical entity defined by its components.
type Vector struct {
	*mat.VecDense
}

// NewVector creates a new Vector instance.
//
// Arguments:
//   - coordinates ([]float64): A slice of float64 values representing the vector's components.
//
// Returns:
//   - *Vector: A pointer to the newly created Vector.
func NewVector(coordinates []float64) *Vector {
	size := len(coordinates)
	return &Vector{
		VecDense: mat.NewVecDense(size, coordinates),
	}
}

// NewVectorTwoPoints creates a vector from two points.
//
// Arguments:
//   - p1 (*Point): The starting point of the vector.
//   - p2 (*Point): The ending point of the vector.
//
// Returns:
//   - *Vector: A pointer to the newly created Vector representing the difference between p2 and p1.
//
// Panics:
//   - If the dimensions of p1 and p2 are not equal.
func NewVectorTwoPoints(p1, p2 *Point) *Vector {
	if p1.Dimension() != p2.Dimension() {
		msg := fmt.Sprintf("Points must have the same dimension: %d != %d", p1.Dimension(), p2.Dimension())
		panic(msg)
	}

	var coordinates []float64
	for i := 0; i < p1.Dimension(); i++ {
		coordinates = append(coordinates, p2.AtVec(i)-p1.AtVec(i))
	}
	size := len(coordinates)
	vec := mat.NewVecDense(size, coordinates)

	return &Vector{
		VecDense: vec,
	}
}

// String returns a string representation of the Vector object.
//
// Returns:
//   - string: A string representation of the Vector.
func (v Vector) String() string {
	return fmt.Sprintf("%v", mat.Formatted(v.VecDense))
}

// Length computes the Euclidean length (2-norm) of the vector.
//
// Returns:
//   - float64: The Euclidean length of the vector.
func (v Vector) Length() float64 {
	return mat.Norm(v.VecDense, 2)
}

// Dimension returns the dimension (number of components) of the vector.
//
// Returns:
//   - int: The dimension of the vector.
func (v Vector) Dimension() int {
	return len(v.RawVector().Data)
}
