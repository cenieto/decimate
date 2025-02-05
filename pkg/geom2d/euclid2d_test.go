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
package geom2d

import (
	"github.com/cenieto/decimate/pkg/primitives"
	"github.com/cenieto/decimate/pkg/testutils"
	"gonum.org/v1/gonum/mat"
	"math"
	"testing"
)

// TestGeom2dInstantiation tests the correct instantiation of a geom2d object.
func TestGeom2dInstantiation(t *testing.T) {
	geom2d := NewEuclid() // Create a new geom2d instance

	dimension := geom2d.Dimension() // Get the dimension of the geometry

	expected := 2 // Expected dimension for a 2D system

	// Verify if the returned dimension matches the expected value
	if dimension != expected {
		t.Errorf("geom2d.Dimension() = %v; want %v", dimension, expected)
	}
}

// TestCrossProduct2D tests the cross product operation for 2D vectors.
// Verifies that the cross product of 2D vectors results in a 3D vector.
// Set of inputs and expected results are read from a CSV file.
func TestCrossProduct2D(t *testing.T) {

	fixtureFile := "../../testdata/geom2d/cross-product.csv"  // Updated fixture file path
	reader, err := testutils.NewCSVFloat64Reader(fixtureFile) // Get the CSV float64 reader
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	lines := reader.ReadLines()

	for values := range lines {
		if values.Err != nil {
			t.Fatalf("Error while reading CSV file: %v", values.Err)
		}

		expected := mat.NewVecDense(3, []float64{
			values.Values[6],
			values.Values[7],
			values.Values[8],
		}) // Expected result for the cross product

		geom := NewEuclid() // Create new geom2d instance

		v1 := primitives.NewVector([]float64{
			values.Values[0],
			values.Values[1],
		}) // First vector for cross product
		v2 := primitives.NewVector([]float64{
			values.Values[3],
			values.Values[4],
		}) // Second vector for cross product

		result := geom.CrossProduct(v1, v2)

		if !mat.EqualApprox(result, expected, 1e-9) {
			t.Errorf("CrossProduct(%v, %v) = %v; want %v", v1, v2, result, expected)
		}
	}

}

// TestCrossProduct2DInvalid tests the CrossProduct method with invalid inputs.
// Verifies that the method panics when given invalid inputs.
func TestCrossProduct2DInvalid(t *testing.T) {
	geom := NewEuclid()

	v1 := primitives.NewVector([]float64{1, 2, 3})
	v2 := primitives.NewVector([]float64{1, 2})

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("CrossProduct(%v, %v) did not panic", v1, v2)
		}
	}()

	geom.CrossProduct(v1, v2)
}

// TestCrossProductNorm tests the cross product norm (magnitude) for 2D vectors.
// Verifies that the cross product of two 2D vectors results in a correct magnitude.
// Set of inputs and expected results are read from a CSV file.
func TestCrossProductNorm(t *testing.T) {
	fixtureFile := "../../testdata/geom2d/cross-product.csv"  // Updated fixture file path
	reader, err := testutils.NewCSVFloat64Reader(fixtureFile) // Get the CSV float64 reader
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	lines := reader.ReadLines()

	for values := range lines {
		if values.Err != nil {
			t.Fatalf("Error while reading CSV file: %v", values.Err)
		}

		expected := math.Abs(values.Values[8]) // Expected result for the magnitude of the cross product

		geom := NewEuclid() // Create new geom2d instance

		v1 := primitives.NewVector([]float64{
			values.Values[0],
			values.Values[1],
		}) // First vector for cross product
		v2 := primitives.NewVector([]float64{
			values.Values[3],
			values.Values[4],
		}) // Second vector for cross product

		// Compute the magnitude of the cross product of v1 and v2
		result := geom.CrossProductNorm(v1, v2)

		// Check if the result is approximately equal to the expected magnitude
		if !mat.EqualApprox(mat.NewVecDense(1, []float64{result}), mat.NewVecDense(1, []float64{expected}), testutils.TestToleranceRelative) {
			t.Errorf("TestCrossProductNorm(%v, %v) = %v; want %v", v1.String(), v2.String(), result, expected)
		}
	}
}

// TestCrossProductNormInvalid tests the CrossProductNorm method with invalid inputs.
// Verifies that the method panics when given invalid inputs.
func TestCrossProductNormInvalid(t *testing.T) {
	geom := NewEuclid()

	v1 := primitives.NewVector([]float64{1, 2, 3})
	v2 := primitives.NewVector([]float64{1, 2})

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("CrossProductNorm(%v, %v) did not panic", v1, v2)
		}
	}()

	geom.CrossProductNorm(v1, v2)
}

// TestDoubleAreaTriangle tests the calculation of the double area of a triangle formed by a point and a line.
// Verifies that the double area of a triangle with vertices in three 2D points is computed correctly.
// Set of inputs and expected results are read from a CSV file.
func TestDoubleAreaTriangle(t *testing.T) {
	fixtureFile := "../../testdata/geom2d/point-line.csv"     // Updated fixture file path
	reader, err := testutils.NewCSVFloat64Reader(fixtureFile) // Get the CSV float64 reader
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	lines := reader.ReadLines()

	for values := range lines {
		if values.Err != nil {
			t.Fatalf("Error while reading CSV file: %v", values.Err)
		}

		geom := NewEuclid()
		point := primitives.NewPoint(
			[]float64{
				values.Values[0],
				values.Values[1],
			},
		)
		point_origin_line := primitives.NewPoint(
			[]float64{
				values.Values[2],
				values.Values[3],
			},
		)
		point_end_line := primitives.NewPoint(
			[]float64{
				values.Values[4],
				values.Values[5],
			},
		)
		line := primitives.NewLine(point_origin_line, point_end_line)

		result := geom.DoubleAreaTriangle(point, line)
		expected := values.Values[6]

		// Check if the result is approximately equal to the expected magnitude
		if !mat.EqualApprox(mat.NewVecDense(1, []float64{result}), mat.NewVecDense(1, []float64{expected}), testutils.TestToleranceRelative) {
			t.Errorf("DoubleAreaTriangle(%v, %v) = %v; want %v", point.String(), line.String(), result, expected)
		}
	}
}

// TestDoubleAreaTriangleInvalid tests the DoubleAreaTriangle method with invalid inputs.
// Verifies that the method panics when given invalid inputs.
func TestDoubleAreaTriangleInvalid(t *testing.T) {
	geom := NewEuclid()

	point := primitives.NewPoint([]float64{1, 2, 3})
	line := primitives.NewLine(primitives.NewPoint([]float64{1, 2}), primitives.NewPoint([]float64{1, 2}))

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DoubleAreaTriangle(%v, %v) did not panic", point, line)
		}
	}()

	geom.DoubleAreaTriangle(point, line)

	point = primitives.NewPoint([]float64{1, 2})
	line = primitives.NewLine(primitives.NewPoint([]float64{1, 2, 3}), primitives.NewPoint([]float64{1, 2}))

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DoubleAreaTriangle(%v, %v) did not panic", point, line)
		}
	}()

	geom.DoubleAreaTriangle(point, line)

	point = primitives.NewPoint([]float64{1, 2})
	line = primitives.NewLine(primitives.NewPoint([]float64{1, 2}), primitives.NewPoint([]float64{1, 2, 3}))

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DoubleAreaTriangle(%v, %v) did not panic", point, line)
		}
	}()

	geom.DoubleAreaTriangle(point, line)
}

// TestDistancePointLine tests the calculation of the shortest distance from a point to a line.
// Verifies that the distance from one 2D point to the line passing through two 2D points is computed correctly.
// Set of inputs and expected results are read from a CSV file.
func TestDistancePointLine(t *testing.T) {
	fixtureFile := "../../testdata/geom2d/point-line.csv"     // Updated fixture file path
	reader, err := testutils.NewCSVFloat64Reader(fixtureFile) // Get the CSV float64 reader
	if err != nil {
		t.Fatalf("Error while opening CSV file: %v", err)
	}

	lines := reader.ReadLines()

	for values := range lines {
		if values.Err != nil {
			t.Fatalf("Error while reading CSV file: %v", values.Err)
		}

		geom := NewEuclid()
		point := primitives.NewPoint(
			[]float64{
				values.Values[0],
				values.Values[1],
			},
		)
		point_origin_line := primitives.NewPoint(
			[]float64{
				values.Values[2],
				values.Values[3],
			},
		)
		point_end_line := primitives.NewPoint(
			[]float64{
				values.Values[4],
				values.Values[5],
			},
		)
		line := primitives.NewLine(point_origin_line, point_end_line)

		result := geom.DistancePointLine(point, line)
		expected := values.Values[7]

		// Check if the result is approximately equal to the expected magnitude
		if !mat.EqualApprox(mat.NewVecDense(1, []float64{result}), mat.NewVecDense(1, []float64{expected}), testutils.TestToleranceRelative) {
			t.Errorf("DistancePointLine(%v, %v) = %v; want %v", point.String(), line.String(), result, expected)
		}
	}
}

// TestDistancePointLineInvalid tests the DistancePointLine method with invalid inputs.
// Verifies that the method panics when given invalid inputs.
func TestDistancePointLineInvalid(t *testing.T) {
	geom := NewEuclid()

	point := primitives.NewPoint([]float64{1, 2, 3})
	line := primitives.NewLine(primitives.NewPoint([]float64{1, 2}), primitives.NewPoint([]float64{1, 2}))

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DistancePointLine(%v, %v) did not panic", point, line)
		}
	}()

	geom.DistancePointLine(point, line)

	point = primitives.NewPoint([]float64{1, 2})
	line = primitives.NewLine(primitives.NewPoint([]float64{1, 2, 3}), primitives.NewPoint([]float64{1, 2}))

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DistancePointLine(%v, %v) did not panic", point, line)
		}
	}()

	geom.DistancePointLine(point, line)

	point = primitives.NewPoint([]float64{1, 2})
	line = primitives.NewLine(primitives.NewPoint([]float64{1, 2}), primitives.NewPoint([]float64{1, 2, 3}))

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DistancePointLine(%v, %v) did not panic", point, line)
		}
	}()

	geom.DistancePointLine(point, line)
}
