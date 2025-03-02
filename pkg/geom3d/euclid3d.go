// Copyright 2025 César Nieto Sánchez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package geom3d

import (
	"fmt"
	"github.com/cenieto/decimate/pkg/decimate"
	"github.com/cenieto/decimate/pkg/primitives"
)

// Euclid3D represents a 3D geometric system.
// It provides the necessary methods to perform 3D geometric operations such as cross product and distance calculations.
type Euclid3D struct {
	Decimate *decimate.Decimate
}

// NewEuclid creates and returns a new instance of Euclid3D.
//
// Returns:
//   - Euclid3D: A new instance of the 3D geometry system.
func NewEuclid() *Euclid3D {
	e := &Euclid3D{}
	e.Decimate = decimate.NewDecimate(*e)
	return e
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
//   - v1 (*primitives.Vector): The first 3D vector to be used in the cross product.
//   - v2 (*primitives.Vector): The second 3D vector to be used in the cross product.
//
// Returns:
//   - *geom3d.Vector3D: A 3D vector where the Z-component represents the cross product of the input 3D vectors.
func (g Euclid3D) CrossProduct(v1, v2 *primitives.Vector) *primitives.Vector {
	errorMsg := ""
	if v1.Dimension() != g.Dimension() {
		errorMsg += fmt.Sprintf("First vector is not 3D, is of dimension %d\n", v1.Dimension())
	}
	if v2.Dimension() != g.Dimension() {
		errorMsg += fmt.Sprintf("Second vector is not 3D, is of dimension %d\n", v2.Dimension())
	}
	if errorMsg != "" {
		errorMsg = fmt.Sprintf("CrossProduct in Euclid3D only accepts 3D vectors.\n %s", errorMsg)
		panic(errorMsg)
	}

	result := []float64{
		v1.At(1, 0)*v2.At(2, 0) - v1.At(2, 0)*v2.At(1, 0),
		-v1.At(0, 0)*v2.At(2, 0) + v1.At(2, 0)*v2.At(0, 0),
		v1.At(0, 0)*v2.At(1, 0) - v1.At(1, 0)*v2.At(0, 0),
	}

	return primitives.NewVector(result)
}

// CrossProductNorm computes the magnitude (norm) of the cross product of two 3D vectors.
// This magnitude corresponds to the area of the parallelogram formed by the vectors.
//
// Parameters:
//   - v1 (*primitives.Vector): The first 3D vector.
//   - v2 (*primitives.Vector): The second 3D vector.
//
// Returns:
//   - float64: The magnitude (norm) of the cross product.
func (g Euclid3D) CrossProductNorm(v1, v2 *primitives.Vector) float64 {
	errorMsg := ""
	if v1.Dimension() != g.Dimension() {
		errorMsg += fmt.Sprintf("First vector is not 3D, is of dimension %d\n", v1.Dimension())
	}
	if v2.Dimension() != g.Dimension() {
		errorMsg += fmt.Sprintf("Second vector is not 3D, is of dimension %d\n", v2.Dimension())
	}
	if errorMsg != "" {
		errorMsg = fmt.Sprintf("CrossProduct in Euclid3D only accepts 3D vectors.\n %s", errorMsg)
		panic(errorMsg)
	}

	crossProduct := g.CrossProduct(v1, v2)
	result := crossProduct.Length()
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
func (g Euclid3D) DoubleAreaTriangle(point *primitives.Point, line *primitives.Line) float64 {

	lineToPoint := primitives.NewVectorTwoPoints(point, line.Point1)
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
func (g Euclid3D) DistancePointLine(point *primitives.Point, line *primitives.Line) float64 {

	numerator := g.DoubleAreaTriangle(point, line)
	denominator := line.VectorDirector().Length()
	return numerator / denominator
}
