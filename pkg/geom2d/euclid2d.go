package geom2d

import (
	"decimator/pkg/geom3d"
)

// Euclid2D represents a 2D geometric system.
// It provides the necessary methods to perform 2D geometric operations such as cross product and others.
type Euclid2D struct {
}

// NewEuclid creates and returns a new instance of Euclid2D.
// This initializes a 2D geometry system ready for performing geometric operations.
//
// Returns:
//   - Euclid2D: A new instance of the 2D geometry system.
func NewEuclid() Euclid2D {
	return Euclid2D{}
}

// Dimension returns the dimension of the geometry system.
// For Euclid2D, this always returns 2, as it represents a 2-dimensional space.
//
// Returns:
//   - int: The dimension of the geometry, which is always 2 for this system.
func (g Euclid2D) Dimension() int {
	return 2
}

// CrossProduct computes the cross product of two 2D vectors and returns the result as a 3D vector.
// The input vectors must have a dimension of 2; otherwise, the function will panic.
// The resulting 3D vector's Z-component represents the cross product result.
//
// Parameters:
//   - v1: The first 2D vector (Vector2D) to be used in the cross product.
//   - v2: The second 2D vector (Vector2D) to be used in the cross product.
//
// Returns:
//   - *geom3d.Vector3D: A 3D vector (Vector3D) representing the cross product,
//     where the Z-component is the result of the 2D cross product.
func (g Euclid2D) CrossProduct(v1, v2 *Vector2D) *geom3d.Vector3D {
	result := geom3d.NewVector(
		0.0,
		0.0,
		v1.At(0, 0)*v2.At(1, 0)-v1.At(1, 0)*v2.At(0, 0),
	)

	return result
}

// CrossProductNorm computes the magnitude (norm) of the cross product of two 2D vectors.
// It returns the magnitude as a float64 value.
//
// Parameters:
//   - v1: The first 2D vector (Vector2D) used in the cross product calculation.
//   - v2: The second 2D vector (Vector2D) used in the cross product calculation.
//
// Returns:
//   - float64: The magnitude (norm) of the cross product result as a scalar value.
func (g Euclid2D) CrossProductNorm(v1, v2 *Vector2D) float64 {
	crossProduct := g.CrossProduct(v1, v2)
	result := crossProduct.VecDense.Norm(2)

	return result
}


func (g Euclid2D) DoubleAreaTriangle(point *Point2D, line *Line2D) float64 {
	lineToPoint := NewVectorTwoPoints(point, line.Point1)
	vectorDirector := line.VectorDirector()
	numerator := g.CrossProductNorm(lineToPoint, vectorDirector)
	return numerator
}

func (g Euclid2D) DistancePointLine(point *Point2D, line *Line2D) float64 {
	numerator := g.DoubleAreaTriangle(point, line)
	denominator := line.VectorDirector().Norm(2)
	return numerator / denominator
}
