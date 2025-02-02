// Copyright 2025 César Nieto Sánchez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package primitives

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// Point represents a point in an n-dimensional space.
//
// This struct wraps a `gonum` VecDense to provide additional functionality
// for representing and manipulating points in n-dimensional geometry.
type Point struct {
	*mat.VecDense
}

// NewPoint creates a new Point instance.
//
// Arguments:
//
//	coordinates ([]float64): A slice of float64 values representing the coordinates of the point.
//
// Returns:
//
//	*Point: A pointer to the newly created Point instance.
//
// Example:
//
//	// Create a 3D point with coordinates (1.0, 2.0, 3.0):
//	point := NewPoint([]float64{1.0, 2.0, 3.0})
func NewPoint(coordinates []float64) *Point {
	size := len(coordinates)
	return &Point{
		VecDense: mat.NewVecDense(size, coordinates),
	}
}

// String returns a string representation of the Point object.
//
// Returns:
//
//	string: A string representation of the Point object in the format "Point{<formatted vector>}".
//
// Example:
//
//	point := NewPoint([]float64{1.0, 2.0})
//	fmt.Println(point) // Output: Point{1.0  2.0}
func (p Point) String() string {
	return fmt.Sprintf("Point{%v}", mat.Formatted(p.VecDense))
}

// Dimension returns the dimension of the point.
//
// This method retrieves the length of the underlying vector, which corresponds
// to the number of dimensions of the point.
//
// Returns:
//
//	int: The number of dimensions of the point.
//
// Example:
//
//	point := NewPoint([]float64{1.0, 2.0, 3.0})
//	fmt.Println(point.Dimension()) // Output: 3
func (v Point) Dimension() int {
	return v.Len()
}
