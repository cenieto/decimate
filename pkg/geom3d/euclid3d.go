package geom3d


// Euclid3D represents a 3D geometric system.
// It provides the necessary methods to perform 3D geometric operations such as cross product and distance calculations.
type Euclid3D struct {
}

// NewEuclid creates and returns a new instance of Euclid3D.
//
// Returns:
//   - Euclid3D: A new instance of the 3D geometry system.
func NewEuclid() Euclid3D {
	return Euclid3D{}
}

// Dimension returns the dimension of the geometry system.
//
// Returns:
//   - int: The dimension of the geometry, which is always 2 for this system.
func (g Euclid3D) Dimension() int {
	return 3
}

// CrossProduct computes the cross product of two 3D vectors and returns the result as a 3D vector.
// The Z-component of the resulting 3D vector represents the scalar cross product of the 3D vectors.
//
// Parameters:
//   - v1 (*Vector3D): The first 3D vector to be used in the cross product.
//   - v2 (*Vector3D): The second 3D vector to be used in the cross product.
//
// Returns:
//   - *Vector3D: A 3D vector where the Z-component represents the cross product of the input 3D vectors.
func (g Euclid3D) CrossProduct(v1, v2 *Vector3D) *Vector3D {
	result := NewVector(
		v1.At(1, 0)*v2.At(2, 0)-v1.At(2, 0)*v2.At(1, 0),
		- v1.At(0, 0)*v2.At(2, 0)+v1.At(2, 0)*v2.At(0, 0),
		v1.At(0, 0)*v2.At(1, 0)-v1.At(1, 0)*v2.At(0, 0),
	)
	return result
}

// CrossProductNorm computes the magnitude (norm) of the cross product of two 3D vectors.
// This magnitude corresponds to the area of the parallelogram formed by the vectors.
//
// Parameters:
//   - v1 (*Vector3D): The first 3D vector.
//   - v2 (*Vector3D): The second 3D vector.
//
// Returns:
//   - float64: The magnitude (norm) of the cross product.
func (g Euclid3D) CrossProductNorm(v1, v2 *Vector3D) float64 {
	crossProduct := g.CrossProduct(v1, v2)
	result := crossProduct.VecDense.Norm(2)
	return result
}

// DoubleAreaTriangle calculates the double of the area of a triangle formed by a point and a line.
// The area is computed using the cross product of the vector from the point to one line endpoint
// and the line's direction vector.
//
// Parameters:
//   - point (*Point3D): The point used to form the triangle.
//   - line (*Line3D): The line forming the base of the triangle.
//
// Returns:
//   - float64: The double of the triangle's area.
func (g Euclid3D) DoubleAreaTriangle(point *Point3D, line *Line3D) float64 {
	lineToPoint := NewVectorTwoPoints(point, line.Point1)
	vectorDirector := line.VectorDirector()
	numerator := g.CrossProductNorm(lineToPoint, vectorDirector)
	return numerator
}

// DistancePointLine computes the shortest distance from a point to a line.
// It divides the double area of the triangle formed by the point and the line
// by the norm of the line's direction vector.
//
// Parameters:
//   - point (*Point3D): The point whose distance to the line is being calculated.
//   - line (*Line3D): The line to which the distance is being measured.
//
// Returns:
//   - float64: The shortest distance from the point to the line.
func (g Euclid3D) DistancePointLine(point *Point3D, line *Line3D) float64 {
	numerator := g.DoubleAreaTriangle(point, line)
	denominator := line.VectorDirector().Norm(2)
	return numerator / denominator
}
